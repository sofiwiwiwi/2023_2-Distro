package main

import (
	"context"
    "bufio"
    "fmt"
    "log"
    "math/rand"
    "os"
    "strconv"
    "time"

	pb "github.com/sofiwiwiwi/2023_1-Distro/tree/bup-develop/Tarea-1/protofiles"
	"google.golang.org/grpc"
)

func main() {
	rand.Seed(time.Now().UnixNano())

    var interesed_users int64

	var f, ar_err = os.Open("parametros_de_inicio.txt")
    if ar_err != nil {
        log.Fatal(ar_err)
    }
    var fileScanner = bufio.NewScanner(f)
    fileScanner.Split(bufio.ScanLines)
 	for fileScanner.Scan() {
        text := fileScanner.Text()
        val, err := strconv.ParseInt(text, 10, 64)
        if err != nil {
            log.Fatalf("Error al analizar la línea %s: %v", text, err)
        }
        interesed_users += val
    }
    while(interesed_users > 0){
	    twtpercent := float64(interesed_users) / 2 * 0.2
	    lower_int := int64(float64(interesed_users)/2 - twtpercent)
	    upper_int := int64(float64(interesed_users)/2 + twtpercent)
	 	SolicitedKeys := rand.Int63n(upper_int-lower_int) + lower_int



	 	//recieve avaible keys
		conn, con_err := grpc.Dial(":50051", grpc.WithInsecure())
		if con_err != nil {
			log.Fatal("Can't connect to server: %v", con_err)
		}

		serviceClient := pb.NewNotifyKeysClient(conn)
		res, l_err := serviceClient.SendKeys(context.Background(), &pb.AvailableKeysReq{
			Id:  1,
			Qty: 5,
		})
		if l_err != nil {
			log.Fatal("Keys are not created" + l_err.Error())
		}
		fmt.Println("Llaves disponibles: " + strconv.FormatInt(res.Qty, 10))
		//rabbitmq queue
		// Connect with Rabbit Queue
		rabbit_conn, rabbit_err := amqp.Dial("amqp://guest:guest@localhost:5672/")

		if rabbit_err != nil {
			fmt.Println(rabbit_err)
			panic(err)
		}

		ch, rab_con_err := rabbit_conn.Channel()
		if rab_con_err != nil {
			fmt.Println(rab_con_err)
		}

		defer ch.Close()

		msgs, send_mq_err := ch.Publish(
			"TestQueue",
			"oceania",   			//chantar nombre de sevidor regional
			SolicitedKeys, 		//numero de keys a solicitar
			true,
			false,
			false,
			false,
			nil,
		)

		if send_mq_err != nil {
			fmt.Println(send_mq_err)
		}

		// notification part
		serviceClient2 := pb.NewFinalNotificationClient(conn)
		res2, err2 := serviceClient2.NotifyRegional(context.Background(), &pb.FinalNotifyRequest{
			NumberOfUsersFailed: int32(1),
		})
		if err2 != nil {
			log.Fatalf("Failed to notify oceania server: %v", err2)
		}

		interesed_users -= (SolicitedKeys - res2.Message)
		log.Printf("oceania server response: %s", res2.Message)
	}
}

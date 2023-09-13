package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"time"

	pb "Tarea-1/protofiles"

	"google.golang.org/grpc"
)

var serv *grpc.Server
var keep_iterating bool = true

type server struct {
	pb.UnimplementedNotifyKeysServer
}

func (s *server) SendKeys(ctx context.Context, req *pb.AvailableKeysReq) (*pb.Empty, error) {
	// Los regionales deberían responder con su nombre solamente
	fmt.Println("Keys received")
	go func() {
		time.Sleep(3)
		serv.Stop()
	}()
	return &pb.Empty{}, nil
}

func (s *server) NotifyContinue(ctx context.Context, req *pb.ContinueServiceReq) (*pb.Empty, error) {
	keep_iterating = req.Continue
	fmt.Println("Continue?: ", keep_iterating)
	go func() {
		time.Sleep(3)
		serv.Stop()
	}()
	return &pb.Empty{}, nil
}

func start_grpc_server() {
	// Establish grpc connection.
	listner, s_err := net.Listen("tcp", ":50054")

	if s_err != nil {
		log.Fatal("Cant create tcp connection")
	}

	serv = grpc.NewServer()
	pb.RegisterNotifyKeysServer(serv, &server{})
	fmt.Println("Watiting for message")

	if err := serv.Serve(listner); err != nil {
		log.Fatal("Can't initialize the server: ", err)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var interesed_users int64

	var f, ar_err = os.Open("clients/parametros_de_inicio.txt")
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
	for interesed_users > 0 && keep_iterating {

		start_grpc_server() // Wait for keys received
		twtpercent := float64(interesed_users) / 2 * 0.2
		lower_int := int64(float64(interesed_users)/2 - twtpercent)
		upper_int := int64(float64(interesed_users)/2 + twtpercent)
		SolicitedKeys := rand.Int63n(upper_int-lower_int) + lower_int
		fmt.Println("Solicited Keys: ", SolicitedKeys)

		start_grpc_server() // Wait for NotifyContinue
		//rabbitmq queue
		// Connect with Rabbit Queue
		// rabbit_conn, rabbit_err := amqp.Dial("amqp://guest:guest@localhost:5672/")

		// if rabbit_err != nil {
		// 	fmt.Println(rabbit_err)
		// 	panic(rabbit_err)
		// }

		// ch, rab_con_err := rabbit_conn.Channel()
		// if rab_con_err != nil {
		// 	fmt.Println(rab_con_err)
		// }

		// defer ch.Close()

		// msgs, send_mq_err := ch.Publish(
		// 	"TestQueue",
		// 	"america",     //chantar nombre de sevidor regional
		// 	SolicitedKeys, //numero de keys a solicitar
		// 	true,
		// 	false,
		// 	false,
		// 	false,
		// 	nil,
		// )

		// if send_mq_err != nil {
		// 	fmt.Println(send_mq_err)
		// }

		// // notification part
		// serviceClient2 := pb.NewFinalNotificationClient(conn)
		// res2, err2 := serviceClient2.NotifyRegional(context.Background(), &pb.FinalNotifyRequest{
		// 	NumberOfUsersFailed: int32(1),
		// })
		// if err2 != nil {
		// 	log.Fatalf("Failed to notify america server: %v", err2)
		// }

		// interesed_users -= (SolicitedKeys - res2.Message)
		// log.Printf("america server response: %s", res2.Message)
	}
}

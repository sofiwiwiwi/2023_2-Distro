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
	"github.com/streadway/amqp"
)


type server struct {
	keysCount int64
	keysReceived chan struct{}
	usersFailed int64
	usersfailedReceived chan struct{}
	pb.UnimplementedNotifyKeysServer
	pb.UnimplementedFinalNotificationServer
}

func (s *server) SendKeys(ctx context.Context, req *pb.AvailableKeysReq) (*pb.Empty, error) {
	fmt.Printf("Keys received: %d", int64(req.Keys))
	s.keysCount += int64(req.Keys)
	s.keysReceived <- struct{}{} // señal de que recibió keys
	return &pb.Empty{}, nil
}

func (s *server) NotifyRegional(ctx context.Context, req *pb.FinalNotifyRequest) (*pb.FinalNotifyResponse, error) {
    s.UsersFailed := req.NumberOfUsersFailed
    message := fmt.Sprintf("%d Usuarios no pudieron acceder a la beta.", numberOfUsersFailed)
    fmt.Println(message)    
	s.usersfailedReceived <- struct{}{} // señal de que recibió keys
    return &pb.Empty{}, nil
}

func main() {
	rand.Seed(time.Now().UnixNano())

    var interesed_users int64

    //lectura inicio
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

	//conexion rbmq
    rabbit_conn, rabbit_err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if rabbit_err != nil {
		fmt.Println(rabbit_err)
		panic(rabbit_err)
	}

	ch, rab_con_err := rabbit_conn.Channel()
	if rab_con_err != nil {
		fmt.Println(rab_con_err)
	}
	defer ch.Close()

	//conexion grpc
	listner, s_err := net.Listen("tcp", ":50051")

	if s_err != nil {
		log.Fatal("Cant create tcp connection")
	}

	serv := grpc.NewServer()
	pb.RegisterNotifyKeysServer(serv, &server{})
	pb.RegisterFinalNotificationServer(serv, &server{})


    while(interesed_users > 0){
	    twtpercent := float64(interesed_users) / 2 * 0.2
	    lower_int := int64(float64(interesed_users)/2 - twtpercent)
	    upper_int := int64(float64(interesed_users)/2 + twtpercent)
	 	SolicitedKeys := rand.Int63n(upper_int-lower_int) + lower_int

	 	//recieve avaible keys ?
		if s_err = serv.Serve(listner); s_err != nil {
			log.Fatal("can't initialize server" + s_err.Error())
		}
		<-serv.keysReceived // wait until the things recieves a message

		//rbmq sending thingy
		queueName := "TestQueue"
		messageBody := fmt.Sprintf("servidor de AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
		send_mq_err := ch.Publish(
			"",
			queueName,
			false,
			false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(messageBody),
			},
		)
		if send_mq_err != nil {
			log.Printf("no se publicó el mensaje : %v", send_mq_err)
		} else {
			log.Printf("se publicó el mensaje : %s", messageBody)
		}
		time.Sleep(time.Second)

		// notification part
		<-serv.usersfailedReceived // wait until the things recieves a message
		NumberUsersFailed = serv.usersFailed
		
		interesed_users -= (SolicitedKeys - NumberUsersFailed)
		log.Printf("Regional server response: %s", NumberUsersFailed)
	}
}
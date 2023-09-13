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
var users_left int32
var interested_users_global int32

type server struct {
	pb.UnimplementedNotifyKeysServer
}

func (s *server) SendKeys(ctx context.Context, req *pb.AvailableKeysReq) (*pb.Empty, error) {
	// Los regionales deberían responder con su nombre solamente
	fmt.Println("Keys received")
	go func() {
		time.Sleep(1 * time.Second)
		serv.Stop()
	}()
	return &pb.Empty{}, nil
}

func (s *server) NotifyContinue(ctx context.Context, req *pb.ContinueServiceReq) (*pb.ContinueServiceReq, error) {
	keep_iterating = req.Continue && interested_users_global > 0
	fmt.Println("Continue?: ", keep_iterating)
	go func() {
		time.Sleep(1 * time.Second)
		serv.Stop()
	}()
	return &pb.ContinueServiceReq{Continue: keep_iterating}, nil
}

func (s *server) UsersNotAdmittedNotify(ctx context.Context, req *pb.UsersNotAdmittedReq) (*pb.Empty, error) {
	users_left = req.Users
	fmt.Println("usuarios sin key?: ", req.Users)
	go func() {
		time.Sleep(1 * time.Second)
		serv.Stop()
	}()
	return &pb.Empty{}, nil
}

func start_grpc_server() {
	// Establish grpc connection.
	listner, s_err := net.Listen("tcp", ":50051")

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

	var interesed_users int32

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
		interesed_users = int32(val)
	}

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
	// queueName := "TestQueue"

	for interesed_users > 0 && keep_iterating {

		start_grpc_server() // Wait for keys received
		twtpercent := float64(interesed_users) / 2 * 0.2
		lower_int := int64(float64(interesed_users)/2 - twtpercent)
		upper_int := int64(float64(interesed_users)/2 + twtpercent)
		SolicitedKeys := rand.Int63n(upper_int-lower_int) + lower_int
		fmt.Println("Solicited Keys: ", SolicitedKeys)
		// messageBody := fmt.Sprintf("oceania,%d", SolicitedKeys)
		// send_mq_err := ch.Publish(
		// 	"",
		// 	queueName,
		// 	false,
		// 	false,
		// 	amqp.Publishing{
		// 		ContentType: "text/plain",
		// 		Body:        []byte(messageBody),
		// 	},
		// )

		// if sen_mq_err != nil {
		// 	log.Printf("no se publicó el mensaje: %v", send_mq_err)
		// } else {
		// 	log.Printf("se publicó el mensaje: %s", messageBody)
		// }d

		start_grpc_server() // Wait for UsersNotAdmittedNotify
		enrolled_users := (int32(SolicitedKeys) - users_left)
		interesed_users -= enrolled_users
		fmt.Println("Se inscribieron", enrolled_users, "personas")
		interested_users_global = interesed_users
		fmt.Println("Quedan", interested_users_global, "personas en espera de cupo")
		// log.Printf("Regional server response: %s", NumberUsersFailed)

		start_grpc_server() // Wait for NotifyContinue
	}
}

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
	keep_iterating = req.Continue && users_left != 0
	fmt.Println("Continue?: ", keep_iterating)
	go func() {
		time.Sleep(1 * time.Second)
		serv.Stop()
	}()
	return &pb.ContinueServiceReq{Continue: keep_iterating}, nil
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

	var interesed_users =0 int64

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
	queueName := "TestQueue"

	for interesed_users > 0 && keep_iterating {

		start_grpc_server() // Wait for keys received
		twtpercent := float64(interesed_users) / 2 * 0.2
		lower_int := int64(float64(interesed_users)/2 - twtpercent)
		upper_int := int64(float64(interesed_users)/2 + twtpercent)
		SolicitedKeys := rand.Int63n(upper_int-lower_int) + lower_int
		fmt.Println("Solicited Keys: ", SolicitedKeys)
		messageBody := fmt.Sprintf("ameica,%d", SolicitedKeys)
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
			log.Printf("no se publicó el mensaje %d: %v", i+1, send_mq_err)
		} else {
			log.Printf("se publicó el mensaje %d: %s", i+1, messageBody)
		}

		start_grpc_server() // Wait for NotifyContinue
	}
}

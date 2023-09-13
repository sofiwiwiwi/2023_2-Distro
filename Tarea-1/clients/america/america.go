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

	"github.com/streadway/amqp"
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
	go func() {
		time.Sleep(2)
		serv.Stop()
	}()
	return &pb.Empty{}, nil
}

func (s *server) NotifyContinue(ctx context.Context, req *pb.ContinueServiceReq) (*pb.ContinueServiceReq, error) {
	keep_iterating = req.Continue && interested_users_global > 0
	go func() {
		time.Sleep(2)
		serv.Stop()
	}()
	return &pb.ContinueServiceReq{Continue: keep_iterating}, nil
}

func (s *server) UsersNotAdmittedNotify(ctx context.Context, req *pb.UsersNotAdmittedReq) (*pb.Empty, error) {
	users_left = req.Users
	go func() {
		time.Sleep(2)
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

	if err := serv.Serve(listner); err != nil {
		log.Fatal("Can't initialize the server: ", err)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

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
		interested_users_global = int32(val)
	}

	rabbitMQServer := os.Getenv("RABBITMQ_SERVER")
	rabbitMQPort := os.Getenv("RABBITMQ_PORT")
	url := fmt.Sprintf("amqp://guest:guest@%s:%s/", rabbitMQServer, rabbitMQPort)
	rabbit_conn, rabbit_err := amqp.Dial(url)

	if rabbit_err != nil {
		log.Fatal(rabbit_err)
	}

	ch, rab_con_err := rabbit_conn.Channel()
	if rab_con_err != nil {
		log.Fatal(rab_con_err)
	}

	defer ch.Close()
	queueName := "TestQueue"

	for interested_users_global > 0 && keep_iterating {

		start_grpc_server() // Wait for keys received
		twtpercent := float64(interested_users_global) / 2 * 0.2
		lower_int := int64(float64(interested_users_global)/2 - twtpercent)
		upper_int := int64(float64(interested_users_global)/2 + twtpercent)
		var SolicitedKeys int64
		if upper_int-lower_int <= 1 {
			SolicitedKeys = upper_int
		} else {
			SolicitedKeys = rand.Int63n(upper_int-lower_int) + lower_int
		}

		fmt.Println("Hay", SolicitedKeys, "personas interesadas en acceder a la beta")
		messageBody := fmt.Sprintf("america,%d", SolicitedKeys)
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
			log.Fatal("no se publicó el mensaje:", send_mq_err)
		}

		start_grpc_server() // Wait for UsersNotAdmittedNotify
		enrolled_users := (int32(SolicitedKeys) - users_left)
		interested_users_global -= enrolled_users
		fmt.Println("Se inscribieron", enrolled_users, "personas")
		fmt.Println("Quedan", interested_users_global, "personas en espera de cupo")

		start_grpc_server() // Wait for NotifyContinue
	}
}

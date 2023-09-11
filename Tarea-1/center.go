package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/streadway/amqp"

	pb "github.com/sofiwiwiwi/2023_1-Distro/tree/bup-develop/Tarea-1/protofiles" // HAY QUE CAMBIAR ESTO AL MAIN CUANDO TODO ESTÉ LISTO
)

// // Escribir las horas

func notify_servers(keys int, register_f *os.File) {
	// Escribir llaves generadas
	// FORMATO:
	// HORA - LLAVES GENERADAS
	// 		Servidor Regional - Llaves Solicitadas - Usuarios Registrados - Usuarios No Registrados
	hr, min, _ := time.Now().Clock()
	var hour_s = fmt.Sprintf("%d : %d", hr, min)
	var _, l_err = register_f.WriteString(hour_s + " - " + strconv.Itoa(keys) + "\n")
	if l_err != nil {
		log.Fatal(l_err)
	}

	// fmt.Println("Waiting for messages...")
	// if s_err = serv.Serve(listner); s_err != nil {
	// 	log.Fatal("can't initialize server" + s_err.Error())
	// }
	return
}

func receive_from_mq(msgs <-chan amqp.Delivery) {
	msg_count := 0
	max_msgs := 4

	consume := make(chan bool)

	func() {
		for d := range msgs {
			fmt.Printf("Mensaje asíncrono: %s de servidor leído\n", d.Body)
			msg_count++

			if msg_count >= max_msgs {
				close(consume)
				break
			}
		}
	}()
}

type server struct {
	pb.UnimplementedNotifyKeysServer
	pb.UnimplementedFinalNotificationServer
}

func (s *server) SendKeys(ctx context.Context, req *pb.AvailableKeysReq) (*pb.AvailableKeysReq, error) {
	// Los regionales deberían responder con su nombre solamente

	fmt.Println("Message received")
	return &pb.AvailableKeysReq{
		Id:  0,
		Qty: 150,
	}, nil
}

func (s *server) NotifyRegional(ctx context.Context, req *pb.FinalNotifyRequest) (*pb.FinalNotifyResponse, error) {
	numberOfUsersFailed := 0 //aca va el numero de usuarios que no pudieron entrar, aun no se como calcularlo
	response := &pb.FinalNotifyResponse{
		Message: fmt.Sprintf("%d Usuarios no pudieron acceder a la beta.", numberOfUsersFailed),
	}

	return response, nil
}

var max_id int64

func main() {
	var register_f, register_err = os.Create("registro_flujo.txt")
	if register_err != nil {
		log.Fatal(register_err)
	}

	defer register_f.Close()

	var f, err = os.Open("parametros_de_inicio.txt")
	if err != nil {
		log.Fatal(err)
	}

	var fileScanner = bufio.NewScanner(f)

	fileScanner.Split(bufio.ScanLines)

	// Read key interval
	fileScanner.Scan()
	var interval = strings.Split(fileScanner.Text(), "-")
	var lower, upper = interval[0], interval[1]
	var lower_int, pl_err = strconv.ParseInt(lower, 0, 0)
	var upper_int, pu_err = strconv.ParseInt(upper, 0, 0)

	// Read Rounds interval
	fileScanner.Scan()
	var rounds = fileScanner.Text()
	var rounds_int, r_err = strconv.ParseInt(rounds, 0, 0)

	f.Close()

	if pl_err != nil {
		log.Fatal(pl_err)
	}
	if pu_err != nil {
		log.Fatal(pu_err)
	}
	if r_err != nil {
		log.Fatal(r_err)
	}

	// Connect with Rabbit Queue
	rabbit_conn, rabbit_err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	if rabbit_err != nil {
		fmt.Println(rabbit_err)
		panic(err)
	}

	ch, err := rabbit_conn.Channel()
	if err != nil {
		fmt.Println(rabbit_err)
	}

	defer ch.Close()

	if err != nil {
		fmt.Println(err)
	}

	msgs, err := ch.Consume(
		"TestQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	// Establish grpc connection.
	// listner, s_err := net.Listen("tcp", ":50051")

	// if s_err != nil {
	// 	log.Fatal("Cant create tcp connection")
	// }

	// serv := grpc.NewServer()
	// pb.RegisterNotifyKeysServer(serv, &server{})
	// pb.RegisterFinalNotificationServer(serv, &server{})
	var i = rounds_int
	for i != 0 {
		fmt.Println("Generación ", i, "/", rounds_int)
		var keys = int(rand.Int63n(upper_int-lower_int) + lower_int)

		notify_servers(keys, register_f)

		upper_int -= 10
		lower_int -= 10

		receive_from_mq(msgs)

		if i > 0 {
			i -= 1
		}
	}
}

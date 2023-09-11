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
	"google.golang.org/grpc"

	pb "github.com/sofiwiwiwi/2023_1-Distro/tree/bup-develop/Tarea-1/protofiles" // HAY QUE CAMBIAR ESTO AL MAIN CUANDO TODO ESTÉ LISTO
)

func receive_from_mq(msgs <-chan amqp.Delivery) {
	msg_count := 0
	max_msgs := 4

	consume := make(chan bool)

	func() { // ver si podemos quitar el chan
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

func send_keys_to_all(keys int, conn_asia *grpc.ClientConn, conn_europe *grpc.ClientConn,
	conn_oceania *grpc.ClientConn, conn_america *grpc.ClientConn) {

	hr, min, _ := time.Now().Clock()
	var hour_s = fmt.Sprintf("%d : %d", hr, min)
	var _, l_err = register_f.WriteString(hour_s + " - " + strconv.Itoa(keys) + "\n")
	if l_err != nil {
		log.Fatal(l_err)
	}
	var clients = [4]*grpc.ClientConn{conn_asia, conn_europe, conn_oceania, conn_america}

	for _, conn := range clients {
		this_client := pb.NewNotifyKeysClient(conn)
		_, l_client_err := this_client.SendKeys(context.Background(), &pb.AvailableKeysReq{
			Id:  1,
			Qty: 5, // reemplazar con keys
		})
		if l_client_err != nil {
			log.Fatal("Couldn't send keys to Asia: ", l_client_err)
		}
	}

	// BORRAR LUEGO DE TESTEAR
	// asia_client := pb.NewNotifyKeysClient(conn_asia)
	// _, l_asia_err := asia_client.SendKeys(context.Background(), &pb.AvailableKeysReq{
	// 	Id:  1,
	// 	Qty: 5,
	// })
	// if l_asia_err != nil {
	// 	log.Fatal("Couldn't send keys to Asia: ", l_asia_err)
	// }

	// europe_client := pb.NewNotifyKeysClient(conn_europe)
	// _, l_europe_err := europe_client.SendKeys(context.Background(), &pb.AvailableKeysReq{
	// 	Id:  1,
	// 	Qty: 5,
	// })
	// if l_europe_err != nil {
	// 	log.Fatal("Couldn't send keys to Europe: ", l_europe_err)
	// }

	// oceania_client := pb.NewNotifyKeysClient(conn_oceania)
	// _, l_oceania_err := oceania_client.SendKeys(context.Background(), &pb.AvailableKeysReq{
	// 	Id:  1,
	// 	Qty: 5,
	// })
	// if l_oceania_err != nil {
	// 	log.Fatal("Couldn't send keys to Asia: ", l_oceania_err)
	// }

	// america_client := pb.NewNotifyKeysClient(conn_america)
	// _, l_america_err := america_client.SendKeys(context.Background(), &pb.AvailableKeysReq{
	// 	Id:  1,
	// 	Qty: 5,
	// })
	// if l_america_err != nil {
	// 	log.Fatal("Couldn't send keys to Asia: ", l_america_err)
	// }
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

var register_f, register_err = os.Create("registro_flujo.txt")

func main() {
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

	conn_asia, con_asia_err := grpc.Dial(":50051", grpc.WithInsecure())
	if con_asia_err != nil {
		log.Fatal("Can't connect to Asia server: ", con_asia_err)
	}
	conn_europe, con_europe_err := grpc.Dial(":50051", grpc.WithInsecure())
	if con_europe_err != nil {
		log.Fatal("Can't connect to Asia server: ", con_europe_err)
	}
	conn_oceania, con_oceania_err := grpc.Dial(":50051", grpc.WithInsecure())
	if con_oceania_err != nil {
		log.Fatal("Can't connect to Asia server: ", con_oceania_err)
	}
	conn_america, con_america_err := grpc.Dial(":50051", grpc.WithInsecure())
	if con_america_err != nil {
		log.Fatal("Can't connect to Asia server: ", con_asia_err)
	}

	var i = rounds_int
	for i != 0 {
		fmt.Println("Generación ", i, "/", rounds_int)
		var keys = int(rand.Int63n(upper_int-lower_int) + lower_int)

		send_keys_to_all(keys, conn_asia, conn_europe, conn_oceania, conn_america)

		upper_int -= 10
		lower_int -= 10

		receive_from_mq(msgs)

		if i > 0 {
			i -= 1
		}
	}
}

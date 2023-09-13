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

	pb "Tarea-1/protofiles" // HAY QUE CAMBIAR ESTO AL MAIN CUANDO TODO ESTÉ LISTO
)

var register_f, register_err = os.Create("registro_flujo.txt")
var conn_asia, conn_america, conn_europe, conn_oceania *grpc.ClientConn
var available [4]bool = [4]bool{false, true, false, false}
var users_left bool = true

func receive_from_mq(msgs <-chan amqp.Delivery) {
	msg_count := 0
	max_msgs := 4

	consume := make(chan bool)

	func() { // ver si podemos quitar el chan
		for d := range msgs {
			fmt.Printf("Mensaje asíncrono: de servidor %s leído\n", d.Body)
			msg_count++

			if msg_count >= max_msgs {
				close(consume)
				break
			}
		}
	}()
}

func connect_to_all() {
	var err error
	if available[0] {
		conn_asia, err = grpc.Dial(":50053", grpc.WithInsecure())
		if err != nil {
			log.Fatal("Can't connect to Asia server: ", err)
		}
	}
	if available[1] {
		conn_america, err = grpc.Dial(":50054", grpc.WithInsecure())
		if err != nil {
			log.Fatal("Can't connect to Asia server: ", err)
		}
	}
	if available[2] {
		conn_europe, err = grpc.Dial(":50052", grpc.WithInsecure())
		if err != nil {
			log.Fatal("Can't connect to Asia server: ", err)
		}
	}
	if available[3] {
		conn_oceania, err = grpc.Dial(":50051", grpc.WithInsecure())
		if err != nil {
			log.Fatal("Can't connect to Asia server: ", err)
		}
	}

	fmt.Println("Connected to everyone correctly")
}

func send_keys_to_all(keys int32) {
	fmt.Println("Enviando llaves a los regionales...")
	hr, min, _ := time.Now().Clock()
	var hour_s = fmt.Sprintf("%d : %d", hr, min)
	var _, l_err = register_f.WriteString(hour_s + " - " + strconv.Itoa(int(keys)) + "\n")
	if l_err != nil {
		log.Fatal(l_err)
	}
	var clients = [4]*grpc.ClientConn{conn_asia, conn_america, conn_europe, conn_oceania}

	for index, conn := range clients {
		if !available[index] {
			continue
		}
		this_client := pb.NewNotifyKeysClient(conn)
		_, l_client_err := this_client.SendKeys(context.Background(), &pb.AvailableKeysReq{
			Keys: keys, // reemplazar con keys
		})
		fmt.Println("Enviado")
		if l_client_err != nil {
			log.Fatal("Couldn't send keys to region: ", l_client_err)
		}
	}
}

func notify_continue_to_all() {
	var clients = [4]*grpc.ClientConn{conn_asia, conn_america, conn_europe, conn_oceania}

	for index, conn := range clients {
		if !available[index] {
			continue
		}
		this_client := pb.NewNotifyKeysClient(conn)
		res, l_client_err := this_client.NotifyContinue(context.Background(), &pb.ContinueServiceReq{
			Continue: users_left, // reemplazar con keys
		})
		if l_client_err != nil {
			log.Fatal("Couldn't send confirm to region: ", l_client_err)
		}
		available[index] = res.Continue
		conn.Close()
	}
}

func notify_users_left_to_all(requested []int32, assigned []int32, servers []string) {
	var clients = [4]*grpc.ClientConn{conn_asia, conn_america, conn_europe, conn_oceania}
	for index, conn := range clients {
		if !available[index] {
			continue
		}
		this_client := pb.NewNotifyKeysClient(conn)
		fmt.Println("Se inscribieron ", assigned[index], "cupos de servidor ", servers[index])
		var not_registered int32 = requested[index] - assigned[index]
		_, l_client_err := this_client.UsersNotAdmittedNotify(context.Background(), &pb.UsersNotAdmittedReq{
			Users: not_registered,
		})
		if l_client_err != nil {
			log.Fatal("Couldn't send users left to region: ", l_client_err)
		}
		conn.Close()
	}
}

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
	// rabbit_conn, rabbit_err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	// if rabbit_err != nil {
	// 	fmt.Println(rabbit_err)
	// 	panic(rabbit_err)
	// }

	// ch, err := rabbit_conn.Channel()
	// if err != nil {
	// 	fmt.Println(rabbit_err)
	// }

	// defer ch.Close()

	// msgs, err := ch.Consume(
	// 	"TestQueue",
	// 	"",
	// 	true,
	// 	false,
	// 	false,
	// 	false,
	// 	nil,
	// )
	// receive_from_mq(msgs)

	var i = rounds_int // debe partir en 1 para el print
	// var ch chan bool
	for i != 0 && users_left {
		if !available[0] && !available[1] && !available[2] && !available[3] {
			break
		}
		fmt.Println("Generación ", i-(i-1), "/", rounds_int)
		var keys = int32(rand.Int63n(upper_int-lower_int) + lower_int)

		// Send keys
		connect_to_all()
		send_keys_to_all(keys)

		time.Sleep(5 * time.Second)

		fmt.Println("Ahora a confirmar si seguimos")

		upper_int -= 10
		lower_int -= 10

		// Receive user peticions
		// receive_from_mq(msgs)

		requested := [4]int32{100, 100, 100, 100}
		assigned := [4]int32{100, 100, 100, 100} // DONT ASSIGN MORE THAN THEY REQUESTED
		servers := [4]string{"asia", "america", "europa", "oceania"}

		connect_to_all()
		notify_users_left_to_all(requested[:], assigned[:], servers[:])
		time.Sleep(5 * time.Second)

		// Notify Continue
		connect_to_all()
		if i > 1 {
			notify_continue_to_all()
		}
		time.Sleep(5 * time.Second)

		if i > 0 {
			i -= 1
		}
	}
	users_left = false
	notify_continue_to_all()
}

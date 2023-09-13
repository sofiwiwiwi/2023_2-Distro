package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/streadway/amqp"
	"google.golang.org/grpc"

	pb "Tarea-1/protofiles"
)

var keys int32

var register_f, register_err = os.Create("registro_flujo.txt")
var conn_asia, conn_america, conn_europe, conn_oceania *grpc.ClientConn
var available [4]bool = [4]bool{true, true, true, true}
var users_left bool = true

var assigned [4]int32
var requested [4]int32
var servers [4]string = [4]string{"asia", "america", "europa", "oceania"}
var order []int

func receive_from_mq(msgs <-chan amqp.Delivery) {
	msgCount := 0
	maxMsgs := 0
	for i := 0; i < 4; i++ {
		if available[i] {
			maxMsgs += 1
		}
	}
	consume := make(chan bool)
	regions := []string{"asia", "america", "europa", "oceania"}
	order = []int{}

	re := regexp.MustCompile(`(\w+),(\d+)`)
	for d := range msgs {
		messageBody := string(d.Body)
		match := re.FindStringSubmatch(messageBody)

		if len(match) >= 3 {
			message := match[1]
			numberStr := match[2]
			number, err := strconv.Atoi(numberStr)
			if err == nil {
				fmt.Println("Mensaje asincrono de servidor", message, "leido")
				for i, region := range regions {
					if message == region {
						requested[i] = int32(number)
						order = append(order, i)
						if keys-requested[i] <= 0 {
							assigned[i] = keys
							keys = 0
						} else {
							assigned[i] = requested[i]
							keys = keys - assigned[i]
						}

						var not_registered int = int(requested[i] - assigned[i])
						var _, l_err = register_f.WriteString("    " + servers[i] + " - " + strconv.Itoa(int(requested[i])) + " - " +
							strconv.Itoa(int(assigned[i])) + " - " + strconv.Itoa(not_registered) + "\n")
						if l_err != nil {
							log.Fatal(l_err)
						}
					}
				}
			}
		}
		msgCount++
		if msgCount >= maxMsgs {
			close(consume)
			break
		}
	}
}

func connect_to_all() {
	var err error
	if available[0] {
		conn_asia, err = grpc.Dial("dist045.inf.santiago.usm.cl:50053", grpc.WithInsecure())
		if err != nil {
			log.Fatal("Can't connect to Asia server: ", err)
		}
	}
	if available[1] {
		conn_america, err = grpc.Dial("dist046.inf.santiago.usm.cl:50054", grpc.WithInsecure())
		if err != nil {
			log.Fatal("Can't connect to America server: ", err)
		}
	}
	if available[2] {
		conn_europe, err = grpc.Dial("dist047.inf.santiago.usm.cl:50052", grpc.WithInsecure())
		if err != nil {
			log.Fatal("Can't connect to Europa server: ", err)
		}
	}
	if available[3] {
		conn_oceania, err = grpc.Dial("dist048.inf.santiago.usm.cl:50051", grpc.WithInsecure())
		if err != nil {
			log.Fatal("Can't connect to Oceania server: ", err)
		}
	}
}

func send_keys_to_all(keys int32) {
	hr, min, ms := time.Now().Clock()
	var hour_s = fmt.Sprintf("%d : %d : %d", hr, min, ms)
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
			Keys: keys,
		})
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
			Continue: users_left,
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
		fmt.Println("Se inscribieron", assigned[index], "cupos de servidor", servers[index])
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

	// Read Rounds
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
	rabbitMQServer := os.Getenv("RABBITMQ_SERVER")
	rabbitMQPort := os.Getenv("RABBITMQ_PORT")
	url := fmt.Sprintf("amqp://guest:guest@%s:%s/", rabbitMQServer, rabbitMQPort)
	rabbit_conn, rabbit_err := amqp.Dial(url)
	if rabbit_err != nil {
		log.Fatal(rabbit_err)
	}

	ch, err := rabbit_conn.Channel()
	if err != nil {
		log.Fatal(rabbit_err)
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

	defer ch.Close()

	var i = rounds_int
	var curr_round = 1
	for i != 0 && users_left {
		if !available[0] && !available[1] && !available[2] && !available[3] {
			break
		}
		fmt.Println("Generación ", curr_round, "/", rounds_int)
		keys = int32(rand.Int63n(upper_int-lower_int) + lower_int)

		// Send keys
		connect_to_all()
		send_keys_to_all(keys)

		time.Sleep(15)

		// Receive user peticions

		receive_from_mq(msgs)

		connect_to_all()
		notify_users_left_to_all(requested[:], assigned[:], servers[:])
		time.Sleep(15)

		// Notify Continue
		connect_to_all()
		if i > 1 || i == -1 {
			notify_continue_to_all()
		}
		time.Sleep(15)

		curr_round += 1
		if i > 0 {
			i -= 1
		}
	}
	users_left = false
	notify_continue_to_all()

	hr, min, ms := time.Now().Clock()
	var hour_s = fmt.Sprintf("%d : %d : %d", hr, min, ms)
	var _, l_err = register_f.WriteString(hour_s + " - " + strconv.Itoa(int(keys)) + "\n")
	if l_err != nil {
		log.Fatal(l_err)
	}
}

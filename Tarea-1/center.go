package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"

	pb "github.com/sofiwiwiwi/2023_1-Distro/tree/bup-develop/Tarea-1/protofiles" // HAY QUE CAMBIAR ESTO AL MAIN CUANDO TODO ESTÉ LISTO

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedNotifyKeysServer
}

func generateID() int64 {
	max_id += 1
	return max_id
}

func (s *server) SendKeys(ctx context.Context, req *pb.AvailableKeysReq) (*pb.AvailableKeysReq, error) {
	fmt.Println("Message received")
	return &pb.AvailableKeysReq{
		Id:  0,
		Qty: 150,
	}, nil
}

var max_id int64

func main() {
	var f, err = os.ReadFile("parametros_de_inicio.txt")
	if err != nil {
		log.Fatal(err)
	}

	var interval = strings.Split(string(f), "-")
	var lower, upper = interval[0], interval[1]
	var lower_int, pl_err = strconv.ParseInt(lower, 0, 0)
	var upper_int, pu_err = strconv.ParseInt(upper, 0, 0)

	if pl_err != nil {
		log.Fatal(pl_err)
	}
	if pu_err != nil {
		log.Fatal(pu_err)
	}

	var keys = rand.Int63n(upper_int-lower_int) + lower_int
	listner, s_err := net.Listen("tcp", ":50051")

	if s_err != nil {
		log.Fatal("Cant create tcp connection")
	}

	serv := grpc.NewServer()
	pb.RegisterNotifyKeysServer(serv, &server{})
	fmt.Println("Waiting for messages...")
	if s_err = serv.Serve(listner); s_err != nil {
		log.Fatal("can't initialize server" + s_err.Error())
	}

	fmt.Println(keys)
}

func notify_servers(keys int64) {

	fmt.Println("INFORMACION SOBRE LAS LLAVES ")
	return
}

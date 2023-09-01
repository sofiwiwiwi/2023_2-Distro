package main

import (
	"context"
	"fmt"
	"log"
	"strconv"

	pb "github.com/sofiwiwiwi/2023_1-Distro/tree/bup-develop/Tarea-1/protofiles"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())

	if err != nil {
		log.Fatal("Can't connect to server")
	}

	serviceClient := pb.NewNotifyKeysClient(conn)

	res, l_err := serviceClient.SendKeys(context.Background(), &pb.AvailableKeysReq{
		Id:  1,
		Qty: 5,
	})

	if l_err != nil {
		log.Fatal("Keys are not created" + l_err.Error())
	}

	fmt.Println("Llaves recibidas: " + strconv.FormatInt(res.Qty, 10))
}

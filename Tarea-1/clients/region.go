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
		log.Fatal("Can't connect to server: %v", err)
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


	// Parte de notificación
	serviceClient2 := pb.NewFinalNotificationClient(conn)
	res2, err2 := serviceClient2.NotifyRegional(context.Background(), &pb.FinalNotifyRequest{
		NumberOfUsersFailed: int32(1),
	})
	if err2 != nil {
		log.Fatalf("Failed to notify regional server: %v", err2)
	}

	log.Printf("Regional server response: %s", res2.Message)
}

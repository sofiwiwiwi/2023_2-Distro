package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	pb "github.com/sofiwiwiwi/2023_2-Distro/tree/main/Tarea-2/protofiles"
)

func main() {
	//conexion a OMS
	conn_OMS, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Can't connect to OMS server: ", err)
	}

	this_client := pb.NewOMSClient(conn_OMS)
	_, l_client_err := this_client.SendNombreEstado(context.Background(), &pb.InfoPersonaContinenteReq{
		Nombre:      "SUS",
		EsInfectado: true,
	})
	if l_client_err != nil {
		log.Fatal("Couldn't send message", l_client_err)
	}
	log.Println("Enviado uwu")
}

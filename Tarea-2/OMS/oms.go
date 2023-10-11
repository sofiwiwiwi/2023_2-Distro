package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/sofiwiwiwi/2023_2-Distro/tree/main/Tarea-2/protofiles"
)

var dataNode1_client pb.DataNodeClient
var dataNode2_client pb.DataNodeClient

type server struct {
	pb.UnimplementedOMSServer
}

func (s *server) SendNombreEstado(ctx context.Context, req *pb.InfoPersonaContinenteReq) (*pb.Empty, error) {
	log.Println(req.Nombre, req.EsInfectado)
	return &pb.Empty{}, nil
}

func (s *server) AskNombres(ctx context.Context, req *pb.InfoPersonasCondicionReq) (*pb.InfoPersonasCondicionResp, error) {
	log.Println(req.EsInfectado)
	return &pb.InfoPersonasCondicionResp{
		Nombres: []string{"JuanCarlos;Bodoque", "SofiGaboJavier;Slay"},
	}, nil
	//HACER UN CHANNEL

}

// Establish grpc connection.
func main() {
	conn_dN1, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Can't connect to OMS server: ", err)
	}

	dataNode1_client = pb.NewDataNodeClient(conn_dN1)

	conn_dN2, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Can't connect to OMS server: ", err)
	}

	dataNode2_client = pb.NewDataNodeClient(conn_dN2)

	listner, s_err := net.Listen("tcp", ":50051")

	if s_err != nil {
		log.Fatal("Cant create tcp connection")
	}

	serv := grpc.NewServer()
	pb.RegisterOMSServer(serv, &server{})

	if err := serv.Serve(listner); err != nil {
		log.Fatal("Can't initialize the server: ", err)
	}
}

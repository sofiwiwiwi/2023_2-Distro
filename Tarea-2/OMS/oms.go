package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/sofiwiwiwi/2023_2-Distro/tree/main/Tarea-2/protofiles"
)

type server struct {
	pb.UnimplementedOMSServer
}

func (s *server) SendNombreEstado(ctx context.Context, req *pb.InfoPersonaContinenteReq) (*pb.Empty, error) {
	log.Println(req.Nombre, req.EsInfectado)
	return &pb.Empty{}, nil
}

// Establish grpc connection.
func main() {
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

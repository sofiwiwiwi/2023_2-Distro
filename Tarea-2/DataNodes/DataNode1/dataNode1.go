package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/sofiwiwiwi/2023_2-Distro/tree/main/Tarea-2/protofiles"
)

type server struct {
	pb.UnimplementedDataNodeServer
}

func (s *server) SendIdEstado(ctx context.Context, req *pb.DatosIdNombreReq) (*pb.Empty, error) {
	log.Println(req.Nombre, req.Id)
	return &pb.Empty{}, nil
}

func (s *server) AskNombreId(ctx context.Context, req *pb.NombrePersonaReq) (*pb.NombrePersonaResp, error) {
	log.Println(req.Id)
	return &pb.NombrePersonaResp{
		Nombre: "Slaaaaay",
	}, nil

}

// Establish grpc connection.
func main() {
	listner, s_err := net.Listen("tcp", ":50051")

	if s_err != nil {
		log.Fatal("Cant create tcp connection")
	}

	serv := grpc.NewServer()
	pb.RegisterDataNodeServer(serv, &server{})

	if err := serv.Serve(listner); err != nil {
		log.Fatal("Can't initialize the server: ", err)
	}
}

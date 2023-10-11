package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"

	"google.golang.org/grpc"

	pb "github.com/sofiwiwiwi/2023_2-Distro/tree/main/Tarea-2/protofiles"
)

var register_f, register_err = os.Create("DATA.txt")

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

func EscribirArchivo(Id int32, Nombre string, Apellido string) {
	//escritura de archivos
	var linea = fmt.Sprintf("%d \t %s \t %s\n", Id, Nombre, Apellido)
	var _, l_err = register_f.WriteString(linea)
	if l_err != nil {
		log.Fatal(l_err)
	}
}

func LeerArchivo(Estado bool) {
	var f, ar_err = os.Open("clients/parametros_de_inicio.txt")
	if ar_err != nil {
		log.Fatal(ar_err)
	}
	var fileScanner = bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)
	fileScanner.Scan()
	for fileScanner.Scan() {
		text := fileScanner.Text()
		splits := strings.Split(text, "\t")

		Id, err := strconv.ParseInt(splits[0], 0, 0)
		Id32 := int32(Id)
		var Nombre string = splits[1]
		var Apellido string = splits[2]

	}
	f.Close()
}

// Establish grpc connection.
func main() {
	defer register_f.Close()
	var _, l_err = register_f.WriteString("ID \t Nombre \t Apellido \n")
	if l_err != nil {
		log.Fatal(l_err)
	}

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

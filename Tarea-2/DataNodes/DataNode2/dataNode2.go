package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"strconv"

	"google.golang.org/grpc"

	pb "github.com/sofiwiwiwi/2023_2-Distro/tree/main/Tarea-2/protofiles"
)

var register_f, register_err = os.Create("DataNodes/DataNode1/DATA.txt")

type server struct {
	pb.UnimplementedDataNodeServer
}

func (s *server) SendIdEstado(ctx context.Context, req *pb.DatosIdNombreReq) (*pb.Empty, error) {
	log.Println(req.Nombre, req.Id)
	NombreCompleto := strings.Split(req.Nombre, ";")
	Nombre := NombreCompleto[0]
	Apellido := NombreCompleto[1]
	EscribirArchivo(req.Id, Nombre, Apellido)
	return &pb.Empty{}, nil
}

func (s *server) AskNombreId(ctx context.Context, req *pb.NombrePersonaReq) (*pb.NombrePersonaResp, error) {
	strId := strconv.Itoa(int(req.Id))
	fmt.Println("recibi una weaaaaaaaaa")
	var nombreRespuestin = LeerArchivo(strId)
	fmt.Println("procesé una weaaaaaaaaa")
	return &pb.NombrePersonaResp{
		Nombre: nombreRespuestin,
	}, nil

}

func EscribirArchivo(Id int32, Nombre string, Apellido string) {
	//escritura de archivos
	var linea = fmt.Sprintf("%d    %s    %s\n", Id, Nombre, Apellido)
	var _, l_err = register_f.WriteString(linea)
	if l_err != nil {
		log.Fatal(l_err)
	}
}

func LeerArchivo(Id string) string { //INUTIL
	var retorno string
	var f, ar_err = os.Open("DataNodes/DataNode1/DATA.txt")
	if ar_err != nil {
		log.Fatal(ar_err)
	}
	var fileScanner = bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)
	fileScanner.Scan()
	for fileScanner.Scan() {
		text := fileScanner.Text()
		splits := strings.Split(text, "    ")
		Id_leido := splits[0]
		if Id_leido == Id {
			var Nombre string = splits[1]
			var Apellido string = splits[2]
			retorno = fmt.Sprintf("%s;%s", Nombre, Apellido)
		}
	}
	f.Close()
	return retorno
}

// Establish grpc connection.
func main() {
	defer register_f.Close()
	var _, l_err = register_f.WriteString("ID    Nombre    Apellido\n")
	if l_err != nil {
		log.Fatal(l_err)
	}

	listner, s_err := net.Listen("tcp", ":50053")

	if s_err != nil {
		log.Fatal("Cant create tcp connection")
	}

	serv := grpc.NewServer()
	pb.RegisterDataNodeServer(serv, &server{})

	if err := serv.Serve(listner); err != nil {
		log.Fatal("Can't initialize the server: ", err)
	}
}

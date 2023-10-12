package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"google.golang.org/grpc"

	pb "github.com/sofiwiwiwi/2023_2-Distro/tree/main/Tarea-2/protofiles"
)

var dataNode1_client pb.DataNodeClient
var dataNode2_client pb.DataNodeClient
var idActual = int32(0)

type server struct {
	pb.UnimplementedOMSServer
}

func (s *server) SendNombreEstado(ctx context.Context, req *pb.InfoPersonaContinenteReq) (*pb.Empty, error) {
	//asumiendo que aca es donde se recibe la info de los continentes, caso contrario mover hasta la estrellita
	log.Println("recibi una peticion de weas")
	log.Println(req.Nombre)
	apellido := strings.Split(req.Nombre, ";")[1]
	log.Println("wi1")
	var dataNodeEscritura int32
	if apellido[0] >= 'A' && apellido[0] <= 'M' {
		dataNodeEscritura = 1
	} else {
		dataNodeEscritura = 2
	}
	log.Println("wi1.5")
	EscribirArchivo(dataNodeEscritura, req.Nombre, req.EsInfectado) // esInfectado=0 cuando ta morto      estrellita
	log.Println("wi2")
	log.Println(req.Nombre, req.EsInfectado)
	log.Println("wi3")
	return &pb.Empty{}, nil
}

func (s *server) AskNombres(ctx context.Context, req *pb.InfoPersonasCondicionReq) (*pb.InfoPersonasCondicionResp, error) {
	log.Println(req.EsInfectado)
	return &pb.InfoPersonasCondicionResp{
		Nombres: []string{"JuanCarlos;Bodoque", "SofiGaboJavier;Slay"},
	}, nil
	//HACER UN CHANNEL

}

// probablemente como el acceso al archivo sera compartido por varios procesos es buena idea revisar lo que me dijo el gabo -> peor caso chantarle una goroutine
func EscribirArchivo(dataNode int32, Nombre string, Estado bool) {
	archivo, err := os.OpenFile("DATA.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer archivo.Close()

	linea := fmt.Sprintf("%d \t dataNode%d \t %s \t %v\n", idActual, dataNode, Nombre, Estado)
	_, err = archivo.WriteString(linea)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("wi1.51")
	//enviar a datanode correspondiente
	if dataNode == 1 {
		_, l_client_err := dataNode1_client.SendIdEstado(context.Background(), &pb.DatosIdNombreReq{
			Id:     idActual,
			Nombre: Nombre,
		})
		if l_client_err != nil {
			log.Fatal("Couldn't send message", l_client_err)
		}
	} else {
		_, l_client_err := dataNode2_client.SendIdEstado(context.Background(), &pb.DatosIdNombreReq{
			Id:     idActual,
			Nombre: Nombre,
		})
		if l_client_err != nil {
			log.Fatal("Couldn't send message", l_client_err)
		}
	}
	log.Println("wi1.52")
	idActual++ //para que sea unicoco
}

// recordar estructura de archivo:
// ID    dataNodex    nombre;apellido    Estado
func LeerArchivo() {
	archivo, err := os.Open("DATA.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer archivo.Close()
	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		linea := scanner.Text()
		a := strings.Split(linea, "\t")
		log.Println(a) //esta wea es pa que no webee por las variables no usadas nomas

		//enviar a ONU, no se como lo haremos aun, por eso no esta hecho :p
		/*
		   _, l_client_err := this_client.SendNombreEstado(context.Background(), &pb.InfoPersonaContinenteReq{
		       Nombre:      Nombre,
		       EsInfectado: Estado,
		   })
		   if l_client_err != nil {
		       log.Fatal("Couldn't send message", l_client_err)
		   }*/

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// Establish grpc connection.
func main() {
	conn_dN1, err := grpc.Dial(":50052", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Can't connect to OMS server: ", err)
	}

	dataNode1_client = pb.NewDataNodeClient(conn_dN1)

	conn_dN2, err := grpc.Dial(":50053", grpc.WithInsecure())
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

package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"sync"

	"google.golang.org/grpc"

	pb "github.com/sofiwiwiwi/2023_2-Distro/tree/main/Tarea-2/protofiles"
)

var dataNode1_client pb.DataNodeClient
var dataNode2_client pb.DataNodeClient

var idActual = int32(0)
var idMu sync.Mutex

var archivo, err = os.Create("OMS/DATA.txt")

type server struct {
	pb.UnimplementedOMSServer
}

func (s *server) SendNombreEstado(ctx context.Context, req *pb.InfoPersonaContinenteReq) (*pb.Empty, error) {
	//asumiendo que aca es donde se recibe la info de los continentes, caso contrario mover hasta la estrellita
	idMu.Lock()
	idActual += 1
	apellido := strings.Split(req.Nombre, ";")[1]
	var dataNodeEscritura int32
	if apellido[0] >= 'A' && apellido[0] <= 'M' {
		dataNodeEscritura = 1
	} else {
		dataNodeEscritura = 2
	}
	EscribirArchivo(dataNodeEscritura, req.Nombre, req.EsInfectado) // esInfectado=0 cuando ta morto      estrellita
	idMu.Unlock()
	return &pb.Empty{}, nil
}

func (s *server) AskNombres(ctx context.Context, req *pb.InfoPersonasCondicionReq) (*pb.InfoPersonasCondicionResp, error) {
	return &pb.InfoPersonasCondicionResp{
		Nombres: []string{"JuanCarlos;Bodoque", "SofiGaboJavier;Slay"},
	}, nil
	//HACER UN CHANNEL

}

// probablemente como el acceso al archivo sera compartido por varios procesos es buena idea revisar lo que me dijo el gabo -> peor caso chantarle una goroutine
func EscribirArchivo(dataNode int32, Nombre string, Estado bool) {
	var esc_estado string
	if Estado {
		esc_estado = "Infectado"
	} else {
		esc_estado = "Muerto"
	}

	linea := fmt.Sprintf("%d \t %d \t %v\n", idActual, dataNode, esc_estado)
	_, err = archivo.WriteString(linea)
	if err != nil {
		log.Fatal(err)
	}
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
}

// recordar estructura de archivo:
// ID    dataNodex    nombre;apellido    Estado
func LeerArchivo() {
	archivo, err := os.Open("OMS/DATA.txt")
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
	if err != nil {
		log.Fatal(err)
	}
	defer archivo.Close()
	archivo.WriteString("ID \t DataNode \t Status\n")

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

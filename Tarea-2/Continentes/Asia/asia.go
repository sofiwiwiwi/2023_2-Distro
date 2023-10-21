package main

import (
	"bufio"
	"context"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"google.golang.org/grpc"

	pb "Tarea-2/protofiles"
)

var this_client pb.OMSClient

func LeerArchivo() {
	var f, ar_err = os.Open("Asia/DATA.txt")
	if ar_err != nil {
		log.Fatal(ar_err)
	}
	var fileScanner = bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)
	fileScanner.Scan()
	i := 1
	for fileScanner.Scan() {
		probabilidad := rand.Float64()
		var isInfectado bool = probabilidad <= 55

		var esc_estado string
		if isInfectado {
			esc_estado = "Infectado"
		} else {
			esc_estado = "Muerto"
		}

		if i > 5 {
			time.Sleep(3 * time.Second)
		}
		text := fileScanner.Text()
		Nombre_formateado := strings.ReplaceAll(text, " ", ";")

		_, l_client_err := this_client.SendNombreEstado(context.Background(), &pb.InfoPersonaContinenteReq{
			Nombre:      Nombre_formateado,
			EsInfectado: isInfectado,
		})
		if l_client_err != nil {
			log.Fatal("Couldn't send message", l_client_err)
		}
		log.Printf("Estado enviado: %s %s\n", text, esc_estado)
		i++
	}
	f.Close()
}

func main() {
	//conexion a OMS
	conn_OMS, err := grpc.Dial("dist048.inf.santiago.usm.cl:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Can't connect to OMS server: ", err)
	}

	this_client = pb.NewOMSClient(conn_OMS)

	LeerArchivo()
}

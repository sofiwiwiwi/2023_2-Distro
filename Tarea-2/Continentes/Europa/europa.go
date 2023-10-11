package main

import (
	"bufio"
	"context"
	"log"
	"math/rand"
	"os"
	"time"

	"google.golang.org/grpc"

	pb "github.com/sofiwiwiwi/2023_2-Distro/tree/main/Tarea-2/protofiles"
)

var this_client pb.OMSClient

func LeerArchivo(Estado string) {
	var f, ar_err = os.Open("DATA.txt")
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

		if i > 5 {
			time.Sleep(3)
		}
		text := fileScanner.Text()
		_, l_client_err := this_client.SendNombreEstado(context.Background(), &pb.InfoPersonaContinenteReq{
			Nombre:      text,
			EsInfectado: isInfectado,
		})
		if l_client_err != nil {
			log.Fatal("Couldn't send message", l_client_err)
		}
		i++
	}
	f.Close()
}

func main() {
	//conexion a OMS
	conn_OMS, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Can't connect to OMS server: ", err)
	}

	this_client = pb.NewOMSClient(conn_OMS)

	log.Println("Enviado uwu")
}

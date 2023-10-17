package main

import (
	"context"
	"log"
	"fmt"

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
	var input string
	input = "INICIAL"
	for (input != "q"){
		fmt.Println("ingrese tipo de consulta: m(muertos), i(infectados), q(terminar programa)")
		fmt.Scanln(&input)
		if(input == "m"){
			ans, l_client_err := this_client.AskNombres(context.Background(), &pb.InfoPersonasCondicionReq{
				EsInfectado: false,
			})
			if l_client_err != nil {
				log.Fatal("Couldn't send message", l_client_err)
			}
			for _, item := range ans.Nombres {
        		fmt.Println(item)
    		}
		}else if(input=="i"){
			ans, l_client_err := this_client.AskNombres(context.Background(), &pb.InfoPersonasCondicionReq{
				EsInfectado: true,
			})
			if l_client_err != nil {
				log.Fatal("Couldn't send message", l_client_err)
			}
			for _, item := range ans.Nombres {
        		fmt.Println(item)
    		}
		}
	}
}

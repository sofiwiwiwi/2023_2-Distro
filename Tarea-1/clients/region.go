package main

import (
	"context"
    "bufio"
    "fmt"
    "log"
    "math/rand"
    "os"
    "strconv"
    "time"

	pb "github.com/sofiwiwiwi/2023_1-Distro/tree/bup-develop/Tarea-1/protofiles"
	"google.golang.org/grpc"
)

func main() {
	rand.Seed(time.Now().UnixNano())

    var interesed_users int64
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())

	var f, errar = os.Open("parametros_de_inicio.txt")
    if errar != nil {
        log.Fatal(errar)
    }
    var fileScanner = bufio.NewScanner(f)
    fileScanner.Split(bufio.ScanLines)
 	for fileScanner.Scan() {
        text := fileScanner.Text()
        val, err := strconv.ParseInt(text, 10, 64)
        if err != nil {
            log.Fatalf("Error al analizar la línea %s: %v", text, err)
        }
        interesed_users += val
    }
    
    twtpercent := float64(interesed_users) / 2 * 0.2
    lower_int := int64(float64(interesed_users)/2 - twtpercent)
    upper_int := int64(float64(interesed_users)/2 + twtpercent)
 	SolicitedKeys := rand.Int63n(upper_int-lower_int) + lower_int

 	fmt.Println("Número aleatorio:", SolicitedKeys)
    fmt.Println("Límite inferior:", lower_int)
    fmt.Println("Límite superior:", upper_int)
/*
el servidor regional deber´a generar un valor entre la
cantidad de usuarios interesados que contiene en su archivo ”parametros de inicio”
dividida en 2 menos el 20% de ese valor y la misma cantidad dividida en 2 m´as el
20% del valor. Es decir, si dentro del archivo ”parametros de inicio” el servidor
regional tiene una cantidad de 300 interesados, deber´a generar un n´umero al azar
entre 120 y 180.*/


	if err != nil {
		log.Fatal("Can't connect to server: %v", err)
	}

	serviceClient := pb.NewNotifyKeysClient(conn)

	res, l_err := serviceClient.SendKeys(context.Background(), &pb.AvailableKeysReq{
		Id:  1,
		Qty: 5,
	})

	if l_err != nil {
		log.Fatal("Keys are not created" + l_err.Error())
	}

	fmt.Println("Llaves recibidas: " + strconv.FormatInt(res.Qty, 10))

	// Parte de notificación
	serviceClient2 := pb.NewFinalNotificationClient(conn)
	res2, err2 := serviceClient2.NotifyRegional(context.Background(), &pb.FinalNotifyRequest{
		NumberOfUsersFailed: int32(1),
	})
	if err2 != nil {
		log.Fatalf("Failed to notify regional server: %v", err2)
	}

	log.Printf("Regional server response: %s", res2.Message)
}

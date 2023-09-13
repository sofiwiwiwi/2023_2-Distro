package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	//Opening a channel to rabbitmq instance over the connection established above
	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
	}

	defer ch.Close()
	//interaction with the instance to declare queues to publish and subscribe

	q, err := ch.QueueDeclare(
		"TestQueue",
		false,
		false,
		false,
		false,
		nil,
	)
	fmt.Println(q)
	defer conn.Close()
}

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

	//print status of queue
	fmt.Println(q)

	if err != nil {
		fmt.Println(err)
	}

	//attempt to publish message to the queue
	err = ch.Publish(
		"",
		"TestQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Hello World"),
		},
	)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfuly published message to queue")

	defer conn.Close()

	fmt.Println("Successfully Connected to our RabbitMQ Instance")
}

package main

import (
	"context"

	"github.com/awgst/go-rabbitmq/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

const rmqUrl = "amqp://guest:guest@localhost:5672/"

func main() {
	// Connect to rabbitmq
	rmqConn, err := rabbitmq.Connect(rmqUrl)
	if err != nil {
		panic(err)
	}
	defer rmqConn.Close()

	// Open a channel
	rmqChannel, err := rmqConn.Channel()
	if err != nil {
		panic(rmqChannel)
	}

	// Publish a message
	err = rmqChannel.PublishWithContext(
		context.Background(),
		"notifications", // exchange
		"",              // routing key
		false,           // mandatory
		false,           // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Hello World!"),
		},
	)
	if err != nil {
		panic(err)
	}

}

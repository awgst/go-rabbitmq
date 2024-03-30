package main

import (
	"fmt"

	"github.com/awgst/go-rabbitmq/rabbitmq"
)

const rmqUrl = "amqp://guest:guest@localhost:5672/"

func main() {
	// Connect to the rabbitmq
	rmqConn, err := rabbitmq.Connect(rmqUrl)
	if err != nil {
		panic(err)
	}

	// Open channel
	rmqChannel, err := rmqConn.Channel()
	if err != nil {
		panic(err)
	}

	// Consume messages from email queue
	messages, err := rmqChannel.Consume(
		"sms", // queue
		"",    // consumer
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		panic(err)
	}

	fmt.Println("Start to listen messages...")
	for msg := range messages {
		fmt.Printf("Message body: %s", msg.Body)
	}
}

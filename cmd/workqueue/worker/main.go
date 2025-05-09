package main

import (
	"bytes"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		panic(err)
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		panic(err)
	}

	running := make(chan struct{})

	go func() {
		for msg := range msgs {
			log.Printf("Received a message: %s, %s", msg.MessageId, msg.Body)
			dots := bytes.Count(msg.Body, []byte("."))
			t := time.Duration(dots)
			time.Sleep(t * time.Second)
			log.Printf("Process message done.: %s", msg.MessageId)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-running
}

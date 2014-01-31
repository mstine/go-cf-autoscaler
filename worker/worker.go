package worker

import (
	"github.com/mstine/go-cf-autoscaler/util"
	"github.com/streadway/amqp"
	"log"
	"os"
)

func Run(uri string) {
	conn, err := amqp.Dial(uri)
	util.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	util.FailOnError(err, "Failed to open a channel")

	defer ch.Close()

	q, err := ch.QueueDeclare(
		"worker", // name
		false,   // durable
		false,   // delete when usused
		false,   // exclusive
		false,   // noWait
		nil,     // arguments
	)
	util.FailOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	util.FailOnError(err, "Failed to register a consumer")

	done := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			//done <- true
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-done
	log.Printf("Done")

	os.Exit(0)
}
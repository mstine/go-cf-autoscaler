package producer

import (
	"fmt"
	"github.com/mstine/go-cf-autoscaler/util"
	"github.com/streadway/amqp"
	"log"
	"os"
)

func Run(uri string) {
	connection, err := amqp.Dial(uri)
	util.FailOnError(err, "Failed to connect to RabbitMQ")
	defer connection.Close()

	channel, err := connection.Channel()
	util.FailOnError(err, "Failed to open a channel")
	defer channel.Close()

	q, err := channel.QueueDeclare(
		"worker", // name
		false,   // durable
		false,   // delete when usused
		false,   // exclusive
		false,   // noWait
		nil,     // arguments
	)
	util.FailOnError(err, "Failed to declare a queue")

	for i := 1;; i++ {
		body := fmt.Sprintf("This is message #%v", i)
		log.Printf("Produced a message: %v", body)

		err = channel.Publish(
		"",         // exchange
		q.Name,     // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType:     "text/plain",
			Body:            []byte(body),
		})
		util.FailOnError(err, "Failed to publish a message")
	}	

	os.Exit(0)
}
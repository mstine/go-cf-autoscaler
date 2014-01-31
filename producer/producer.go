package producer

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"os"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func Run() {
	connection, err := amqp.Dial("amqp://localhost:5672")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer connection.Close()

	channel, err := connection.Channel()
	failOnError(err, "Failed to open a channel")
	defer channel.Close()

	q, err := channel.QueueDeclare(
		"worker", // name
		false,   // durable
		false,   // delete when usused
		false,   // exclusive
		false,   // noWait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

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
		failOnError(err, "Failed to publish a message")
	}	

	os.Exit(0)
}
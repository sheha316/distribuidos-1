package main

import (
	"log"
	"encoding/csv"
	"fmt"
	"os"
	"io"
	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://test:test@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()
	q, err := ch.QueueDeclare(
		"finances", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
	    checkError("Cannot create file", err)
			Finances(d.Body)
		}
	}()
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

}
	func Finances(body string) {
		f, err := os.OpenFile("../storage/finanzas/result.csv", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			log.Fatalln("Couldn't open the csv file", err)
		}
		defer f.Close()
		var data [][]string
	  data = append(data, []string{body})
	  w := csv.NewWriter(f)
		w.WriteAll(data)
		f.Close()
}

	func checkError(message string, err error) {
	    if err != nil {
	        log.Fatal(message, err)
	    }
	}

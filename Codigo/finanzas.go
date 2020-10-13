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
	conn, err := amqp.Dial("amqp://test:test@dist93:5672/")
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
			file, err := os.Create("result.csv")
	    checkError("Cannot create file", err)
	    defer file.Close()

	    writer := csv.NewWriter(file)
	    defer writer.Flush()

			var data = [][]string{{"Line1", "Hello Readers of"}, {"Line2", "golangcode.com"}}

	    for _, value := range data {
	        err := writer.Write(value)
	        checkError("Cannot write to file", err)
	    }
			Finances()

		}

}()
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever


	func Finances() {
		// Open the file
		csvfile, err := os.Open("result.csv")
		if err != nil {
			log.Fatalln("Couldn't open the csv file", err)
		}

		// Parse the file
		r := csv.NewReader(csvfile)
		//r := csv.NewReader(bufio.NewReader(csvfile))

		// Iterate through the records
		for {
			// Read each record from csv
			record, err := r.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("cosa: %s \n", record)
		}
	}

	func checkError(message string, err error) {
	    if err != nil {
	        log.Fatal(message, err)
	    }
	}

package main

import (
	"log"
	"encoding/csv"
	"os"
	"strconv"
	"github.com/streadway/amqp"
	"encoding/json"
)

type finanzas struct{
  Id string
  Tipo string
  Valor string
  Intentos string
  Fecha string
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	var balance float32
	balance=0
	os.Remove("../storage/finanzas/result.csv")
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
		var aux finanzas
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
	    checkError("Cannot create file", err)
			json.Unmarshal(d.Body,&aux)
			balance+=Finances(aux)
		}
	}()
	log.Printf("Balance Total: $%d dignipesos",balance)
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

}
	func Finances(body finanzas)(float32) {
		id:=body.Id
		tipo:=body.Tipo
		valor:=body.Valor
		intentos:=body.Intentos
		fech:=body.Fecha


		var ganancia float32
		var perdida int
		var aux int
		perdida,_=strconv.Atoi(intentos)
		aux,_=strconv.Atoi(valor)
		ganancia=0
		if(fech!="0"){
			perdida-=1
			ganancia=float32(aux)

		}else{
			if(tipo=="retail"){
				ganancia=float32(aux)
			}else if(tipo=="prioritario"){
				ganancia=float32(aux)
				ganancia*=float32(0.3)
			}
		}
		perdida*=10
		ganancia-=float32(perdida)
		f, err := os.OpenFile("../storage/finanzas/result.csv", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			log.Fatalln("Couldn't open the csv file", err)
		}
		defer f.Close()
		var data [][]string
	  data = append(data, []string{id+","+tipo+","+valor+","+intentos+","+fech+","+strconv.FormatFloat(ganancia, 'E', -1, 64)})
	  w := csv.NewWriter(f)
		w.WriteAll(data)
		f.Close()
		return ganancia
}

	func checkError(message string, err error) {
	    if err != nil {
	        log.Fatal(message, err)
	    }
	}

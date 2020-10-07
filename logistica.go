package main


import (
	"fmt"
	"log"
	"net"
	"google.golang.org/grpc"
)

type Logistica struct{
	Tipo string
}

type registro_logistica struct{
	timestamp string
	id int
	tipo string
	nombre string
	valor int
	origen string
	destino string
	seguimento int
}

type paquete_logistica struct{
	id int
	seguimento int
	tipo string
	valor int
	intentos int
	estado string
}

func main() {
	lis,err := net.Listen("tcp",":9000")
	if err!=nil{
		log.Fatalf("No encontro puerto 9000 disponible: %v",err)
	}
	grcpserver := grpc.NewServer()
	if err:= grcpserver.Serve(lis); err!= nil{
		log.Fatalf("fallo levantar el servidor de grpc")
	}

   fmt.Println("Hello, World!")
}

package main


import ("fmt"	"time"	"log"	"net"	"google.golang.org/grpc") 

type paquete_loistica struct{
	id int
	seguimento int
	tipo string
	valor int
	intentos int
	estado string
}

type Camion struct{
	Tipo string
	paquete1 paquete_loistica
	paquete2 paquete_loistica
	
}

func main() {
	
	fmt.Println("Hello, World!")
}
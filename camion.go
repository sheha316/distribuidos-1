package main


import (
	"fmt"
	"log"
	"google.golang.org/grpc"
)

// type paquete_logistica struct{
// 	id int
// 	seguimento int
// 	tipo string
// 	valor int
// 	intentos int
// 	estado string
// }
//
// type Camion struct{
// 	Tipo string
// 	paquete1 paquete_logistica
// 	paquete2 paquete_logistica
//
// }

func main() {
	fmt.Println("Hello, World!")
	// var conn *grpc.ClientConn
	conn, err := grpc.Dial("dist93:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	// defer conn.Close()


}

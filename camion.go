package main


import ("fmt"	"time"	"log"	"net"	"google.golang.org/grpc")

type paquete_logistica struct{
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
	var conn *grpc.ClientConn
  conn, err := grpc.Dial(":9000", grpc.WithInsecure())
  if err != nil {
    log.Fatalf("could not connect: %s", err)
  }
  defer conn.Close()
	fmt.Println("Hello, World!")
}

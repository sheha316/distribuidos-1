package main


import (
	"fmt"
	"log"
	"net"
	"context"
	"google.golang.org/grpc"
)

// type Logistica struct{
// 	Tipo string
// }
//
// type registro_logistica struct{
// 	timestamp string
// 	id int
// 	tipo string
// 	nombre string
// 	valor int
// 	origen string
// 	destino string
// 	seguimento int
// }
//
// type paquete_logistica struct{
// 	id int
// 	seguimento int
// 	tipo string
// 	valor int
// 	intentos int
// 	estado string
// }

type logisticaServer struct {

}

// test = {
// 	timestamp = "hola",
// 	id = 1,
// 	tipo = "hola",
// 	nombre = "hola",
// 	valor = 4392,
// 	origen = "ncjds",
// 	destino = "cueic",
// 	seguimento = 4382
// }

func (s *logisticaServer) GetRegistro(ctx context.Context, Camion *proto.Camion) (registro *proto.Registro_logistica, error) {
	camion := Camion.GetId()
	palabra := "hola"
	return palabra
	// return test, nil
}

func main() {
	fmt.Println("Hello, World!")
	lis,err := net.Listen("tcp",":9000")
	if err!=nil{
		log.Fatalf("No encontro puerto 9000 disponible: %v",err)
	}
	grcpserver := grpc.NewServer()
	if err:= grcpserver.Serve(lis); err!= nil{
		log.Fatalf("fallo levantar el servidor de grpc")
	}

}

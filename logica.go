package main
import (
  "log"
  "github.com/sheha316/distribuidos-1/comms"
)

type envio struct{
  Id_paquete string
  Estado string
  Id_camion string
  Seguimiento int
  Tipo string
  Valor int
  Intentos int
  Uso string
}
func main() {
  var envios_s [6]envio
  for i:=0;i<6;i++{
    envios_s[i].Uso="0"
  }
  lis, err := net.Listen("tcp", ":9000")
  if err != nil {
    log.Fatalf("failed to listen: %v", err)
  }

  s := comms.Server{}

  grpcServer := grpc.NewServer()

  comms.RegisterCommsServer(grpcServer, &s)
  if err := grpcServer.Serve(lis); err != nil {
    log.Fatalf("failed to serve: %s", err)
  }
}

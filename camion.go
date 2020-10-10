package main
import (
  "log"
  "github.com/sheha316/distribuidos-1/comms"
  "golang.org/x/net/context"
  "google.golang.org/grpc"
)
type paquete_info struct{
  Id string
  Tipo string
  Valor int
  Tienda string
  Destino string
  Intentos int
}
type Camion struct{
  Tipo string
  Paquete_inf paquete_info
}
func request_paquete(conn *grpc.ClientConn, kamion *Camion){
  c := comms.NewCommsClient(conn)
  response, _ := c.SolicitarPaquete(context.Background(), &comms.Request_SolicitarPaquete{Tipo: kamion.Tipo})
  if(int(response.Valor)!=-1){
    kamion.Paquete_inf=paquete_info{Id:response.Id,Tipo:response.Tipo,Valor:int(response.Valor),
    Tienda:response.Tienda,Destino:response.Destino,Intentos:0}
  }
}

func main() {
  /*camion_1:=&Camion{
    Tipo: "retail",}
  camion_2:=&Camion{
    Tipo: "retail",}*/
  camion_3:=&Camion{
    Tipo: "normal",}
  var conn *grpc.ClientConn
  conn, err := grpc.Dial("dist93:9000", grpc.WithInsecure())
  if err != nil {
    log.Fatalf("did not connect: %+v", err)
  }
  defer conn.Close()
  request_paquete(conn,camion_3)
  log.Printf("Response from server: %s",camion_3.Paquete_inf)
}

package main
import (
  "log"
  "github.com/sheha316/distribuidos-1/Codigo/comms"
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
  Id string
  Paquete_inf1 paquete_info
  Paquete_inf2 paquete_info
  Paquetes int
  Estado int
}

func request_paquete_1(conn *grpc.ClientConn, kamion *Camion){
  c := comms.NewCommsClient(conn)
  response, _ := c.SolicitarPaquete(context.Background(), &comms.Request_SolicitarPaquete{Tipo: kamion.Tipo})
  log.Printf("Response from server: %+v",response)
  if(int(response.Valor)!=-1){
    kamion.Paquete_inf1=paquete_info{Id:response.Id,Tipo:response.Tipo,Valor:int(response.Valor),
    Tienda:response.Tienda,Destino:response.Destino,Intentos:0}
  }
}
func request_paquete_2(conn *grpc.ClientConn, kamion *Camion){
  c := comms.NewCommsClient(conn)
  response, _ := c.SolicitarPaquete(context.Background(), &comms.Request_SolicitarPaquete{Tipo: kamion.Tipo,Id:kamion.Id})
  log.Printf("Response from server: %+v",response)
  if(int(response.Valor)!=-1){
    kamion.Paquete_inf2=paquete_info{Id:response.Id,Tipo:response.Tipo,Valor:int(response.Valor),
    Tienda:response.Tienda,Destino:response.Destino,Intentos:0}
  }
}

func main() {
  /*camion_1:=&Camion{
    Tipo: "retail",}
  camion_2:=&Camion{
    Tipo: "retail",}*/
  camion_3:=&Camion{
    Tipo: "normal",Paquetes:0,Estado:0,Id:"3"}
  var conn *grpc.ClientConn
  conn, err := grpc.Dial("dist93:9000", grpc.WithInsecure())
  if err != nil {
    log.Fatalf("did not connect: %s", err)
  }
  defer conn.Close()

  request_paquete_1(conn,camion_3)
}

package main
import (
  "log"
  "github.com/sheha316/distribuidos-1/Codigo/comms"
  "golang.org/x/net/context"
  "google.golang.org/grpc"
  "strconv"
  "math/rand"
  "time"
  "fmt"
)

func read_and_request_pymes(conn *grpc.ClientConn,id int)(int){
  c := comms.NewCommsClient(conn)
  var tiendas [3]string
  tiendas[0]="re-play"
  tiendas[1]="hamazon"
  tiendas[2]="batalla.net"
  var destino [3]string
  destino[0]="mi casa"
  destino[1]="tu casa"
  destino[2]="nuestra casa"
  var producto [3]string
  producto[0]="dulce"
  producto[1]="pelicula"
  producto[2]="gorro"
  response, err := c.CrearOrdenPyme(context.Background(),&comms.Request_CrearOrdenPyme{
    Id:strconv.Itoa(id),
    Producto:producto[rand.Intn(3)],
    Valor:int32(rand.Intn(30)),
    Tienda:tiendas[rand.Intn(3)],
    Destino:destino[rand.Intn(3)],
    Prioritario:int32(rand.Intn(2))})
  if err != nil {
    log.Fatalf("Error when calling SayHello: %s", err)
  }
  log.Printf("Numero de seguimento: %d", int(response.Seguimiento))
  return int(response.Seguimiento)
}

func read_and_request_retail(conn *grpc.ClientConn,id int)(int){
  c := comms.NewCommsClient(conn)
  var tiendas [3]string
  tiendas[0]="re-play"
  tiendas[1]="hamazon"
  tiendas[2]="batalla.net"
  var destino [3]string
  destino[0]="mi casa"
  destino[1]="tu casa"
  destino[2]="nuestra casa"
  var producto [3]string
  producto[0]="dulce"
  producto[1]="pelicula"
  producto[2]="gorro"
  response, err := c.CrearOrdenRetail(context.Background(),&comms.Request_CrearOrdenRetail{
    Id:strconv.Itoa(id),
    Producto:producto[rand.Intn(3)],
    Valor:int32(rand.Intn(30)),
    Tienda:tiendas[rand.Intn(3)],
    Destino:destino[rand.Intn(3)],})
  if err != nil {
    log.Fatalf("Error when calling SayHello: %s", err)
  }
  log.Printf("Numero de seguimento: %d", int(response.Seguimiento))
  return int(response.Seguimiento)
}

func limpiar(conn *grpc.ClientConn){
  c := comms.NewCommsClient(conn)
  c.LimpiarRegistros(context.Background(),&comms.Dummy{})
}

func send_seguimento(conn *grpc.ClientConn,codigo int){
  c := comms.NewCommsClient(conn)
  response, err := c.Seguimiento(context.Background(),&comms.Request_Seguimiento{Seguimiento:int32(codigo)})
  if err != nil {
    log.Fatalf("Error when calling SayHello: %s", err)
  }
  log.Printf("Estado del paquete "+strconv.Itoa(codigo)+" : %s",(response.Estado))

}

func main() {
  var conn *grpc.ClientConn
  conn, err := grpc.Dial("dist93:9001", grpc.WithInsecure())
  if err != nil {
    log.Fatalf("did not connect: %s", err)
  }
  defer conn.Close()
  var codigos [100]int
  var entregados int
  var opcion int
  var timex int
  limpiar(conn)

  entregados=0
  codigos[entregados]=read_and_request_pymes(conn,entregados)
  entregados++
  log.Printf("Tiempo entre inputs del cliente")
  fmt.Scanln(&timex)
  for{
    time.Sleep(time.Duration(timex) * time.Second)
    opcion=rand.Intn(3)
    switch opcion {
    case 0:
      if(entregados<100){
        codigos[entregados]=read_and_request_pymes(conn,entregados)
        entregados++
      }
    case 1:
      if(entregados<100){
        codigos[entregados]=read_and_request_retail(conn,entregados)
        entregados++
      }
    case 2:
      send_seguimento(conn,codigos[rand.Intn(entregados)])
    }
  }
}

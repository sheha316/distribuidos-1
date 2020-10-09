package main
import (
  "log"
  "github.com/sheha316/distribuidos-1/comms"
  "golang.org/x/net/context"
  "google.golang.org/grpc"
  "os"
  "io"
  "encoding/csv"
)

type Pedido_pymes struct{
  id string 'json:"id"'
  producto string 'json:"producto"'
  valor int 'json:"valor"'
  tienda string 'json:"tienda"'
  destino string 'json:"destino"'
  prioritario int'json:"prioritario"'
}
func read_and_request_pymes(c *commsClient){
  csvFile,_:=os.Open("Prueba/pymes.csv")
  reader := csv.NewReader(bufio.NewReader(csvFile))
  var pedido_pymes []Pedido_pymes
  for{
    line,error :=reader.Read()
    if error==io.EOF{
      break
    }else if error!=nil{
      log.Fatal(error)
    }
    pedido_pymes=append(pedido_pymes,Pedido_pymes{
      id:line[0],
      producto:line[1],
      valor:line[2],
      tienda:line[3],
      destino:line[4],
      prioritario:line[5],
    })
  }
  for i:=0; i<len(pedido_pymes);i++
  {
    response, err := c.Seguimiento(context.Background(), &comms.Request_Seguimiento{Seguimiento: 1})
    response, err := c.CrearOrdenPyme(context.Background(),&comms.Request_CrearOrdenPyme{
      Id:pedido_pymes[i].id,
      Producto:pedido_pymes[i].producto,
      Valor:pedido_pymes[i].valor,
      Tienda:pedido_pymes[i].tienda,
      Destino:pedido_pymes[i].destino,
      Prioritario:pedido_pymes[i].prioritario,})
    if err != nil {
      log.Fatalf("Error when calling SayHello: %s", err)
    }
    log.Printf("Response from server: %s", response.Seguimiento)
  }
}


func main() {
  var conn *grpc.ClientConn
  conn, err := grpc.Dial("dist93:9000", grpc.WithInsecure())
  if err != nil {
    log.Fatalf("did not connect: %s", err)
  }
  defer conn.Close()
  c := comms.NewCommsClient(conn)
  read_and_request_pymes()

  response, err := c.Seguimiento(context.Background(), &comms.Request_Seguimiento{Seguimiento: 1})
  if err != nil {
    log.Fatalf("Error when calling SayHello: %s", err)
  }
  log.Printf("Response from server: %s", response.Estado)
}

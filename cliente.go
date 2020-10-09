package main
import (
  "bufio"
  "log"
  "github.com/sheha316/distribuidos-1/comms"
  "golang.org/x/net/context"
  "google.golang.org/grpc"
  "os"
  "io"
  "encoding/csv"
  "strconv"
)

type Pedido_pymes struct{
  Id string
  Producto string
  Valor int
  Tienda string
  Destino string
  Prioritario int
}
func read_and_request_pymes(c grpc.ClientConnInterface){
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
    line[2]:_=strconv.Atoi(line[2])
    line[5]:_=strconv.Atoi(line[5])
    pedido_pymes=append(pedido_pymes,Pedido_pymes{
      Id:line[0],
      Producto:line[1],
      Valor:line[2],
      Tienda:line[3],
      Destino:line[4],
      Prioritario:line[5],
    })
  }
  for i:=0; i<len(pedido_pymes);i++{
    response, err := c.CrearOrdenPyme(context.Background(),&comms.Request_CrearOrdenPyme{
      Id:pedido_pymes[i].Id,
      Producto:pedido_pymes[i].Producto,
      Valor:pedido_pymes[i].Valor,
      Tienda:pedido_pymes[i].Tienda,
      Destino:pedido_pymes[i].Destino,
      Prioritario:pedido_pymes[i].Prioritario,})
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
  read_and_request_pymes(c)

  response, err := c.Seguimiento(context.Background(), &comms.Request_Seguimiento{Seguimiento: 1})
  if err != nil {
    log.Fatalf("Error when calling SayHello: %s", err)
  }
  log.Printf("Response from server: %s", response.Estado)
}

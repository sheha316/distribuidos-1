package main
import (
  "bufio"
  "log"
  "github.com/sheha316/distribuidos-1/comms"
  "golang.org/x/net/context"
  "google.golang.org/grpc"
  "os"
  "io"
  "fmt"
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

type Pedido_retail struct{
  Id string
  Producto string
  Valor int
  Tienda string
  Destino string
}
func read_and_request_pymes(conn *grpc.ClientConn){
  c := comms.NewCommsClient(conn)
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
    aux1,_:=strconv.Atoi(line[2])
    aux2,_:=strconv.Atoi(line[5])
    pedido_pymes=append(pedido_pymes,Pedido_pymes{
      Id:line[0],
      Producto:line[1],
      Valor:aux1,
      Tienda:line[3],
      Destino:line[4],
      Prioritario:aux2,
    })
  }
  for i:=0; i<len(pedido_pymes);i++{
    response, err := c.CrearOrdenPyme(context.Background(),&comms.Request_CrearOrdenPyme{
      Id:pedido_pymes[i].Id,
      Producto:pedido_pymes[i].Producto,
      Valor:int32(pedido_pymes[i].Valor),
      Tienda:pedido_pymes[i].Tienda,
      Destino:pedido_pymes[i].Destino,
      Prioritario:int32(pedido_pymes[i].Prioritario),})
    if err != nil {
      log.Fatalf("Error when calling SayHello: %s", err)
    }
    log.Printf("Response from server: %d", int(response.Seguimiento))
  }
}
func read_and_request_retail(conn *grpc.ClientConn){
  c := comms.NewCommsClient(conn)
  csvFile,_:=os.Open("Prueba/retail.csv")
  reader := csv.NewReader(bufio.NewReader(csvFile))
  var pedido_retail []Pedido_retail
  for{
    line,error :=reader.Read()
    if error==io.EOF{
      break
    }else if error!=nil{
      log.Fatal(error)
    }
    aux1,_:=strconv.Atoi(line[2])
    pedido_retail=append(pedido_retail,Pedido_retail{
      Id:line[0],
      Producto:line[1],
      Valor:aux1,
      Tienda:line[3],
      Destino:line[4],
    })
  }
  for i:=0; i<len(pedido_retail);i++{
    response, err := c.CrearOrdenRetail(context.Background(),&comms.Request_CrearOrdenRetail{
      Id:pedido_retail[i].Id,
      Producto:pedido_retail[i].Producto,
      Valor:int32(pedido_retail[i].Valor),
      Tienda:pedido_retail[i].Tienda,
      Destino:pedido_retail[i].Destino,})
    if err != nil {
      log.Fatalf("Error when calling SayHello: %s", err)
    }
    log.Printf("Response from server: %d", int(response.Seguimiento))
  }
}


func send_seguimento(conn *grpc.ClientConn){
  c := comms.NewCommsClient(conn)
  log.Printf("Ingrese Numero de Seguimento por favor")
  fmt.Scanln(&input_us)
  response, err := c.Seguimiento(context.Background(),&comms.Request_Seguimiento{Seguimiento:strconv.Atoi(input_us)})
  if err != nil {
    log.Fatalf("Error when calling SayHello: %s", err)
  }
  log.Printf("Response from server: %s", (response.Seguimiento))

}

func main() {
  var conn *grpc.ClientConn
  conn, err := grpc.Dial("dist93:9000", grpc.WithInsecure())
  if err != nil {
    log.Fatalf("did not connect: %s", err)
  }
  defer conn.Close()
  var input_us string
  input_us=""

  for input_us!="0"{
    log.Printf("Bienvenido! ingrese el numero de la opcion que desea")
    log.Printf("1-Hacer pedidos pymes")
    log.Printf("2-Hacer pedidos retail")
    log.Printf("3-Hacer Todos los pedidos")
    log.Printf("4-Hacer Seguimento")
    log.Printf("0-exit")
    fmt.Scanln(&input_us)
    switch  input_us{
      case "1":
        read_and_request_pymes(conn)
  	  case "2":
        read_and_request_retail(conn)
      case "3":
        read_and_request_pymes(conn)
        read_and_request_retail(conn)
      case "4":
        send_seguimento(conn)
  	  default:
  		// freebsd, openbsd,
  	}
  }


  response, err := c.Seguimiento(context.Background(), &comms.Request_Seguimiento{Seguimiento: 1})
  if err != nil {
    log.Fatalf("Error when calling SayHello: %s", err)
  }
  log.Printf("Response from server: %s", response.Estado)
}

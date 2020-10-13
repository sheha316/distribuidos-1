package main
import (
  "os"
  "log"
  "encoding/csv"
  "bufio"
  "io"
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
  csvfilez ,_:= os.Open("../Pedidos/pymes.csv")
  readerz := csv.NewReader(bufio.NewReader(csvfilez))
  var line
  for i:=0; true ;i++{
    line,error :=readerz.Read()
    if error==io.EOF{
      return -1
    }else if error!=nil{
        log.Fatal(error)
        continue
    }
    if i==id{
      aux1,_:=strconv.Atoi(line[2])
      aux2,_:=strconv.Atoi(line[5])
      response, err := c.CrearOrdenPyme(context.Background(),&comms.Request_CrearOrdenPyme{
        Id:line[0],
        Producto:line[1],
        Valor:int32(aux1),
        Tienda:line[3],
        Destino:line[4],
        Prioritario:int32(aux2)})
      if err != nil {
        log.Fatalf("Error when calling SayHello: %s", err)
      }
      break
    }
  }

  csvfilez.Close()
  log.Printf("Numero de seguimento: %d", int(response.Seguimiento))
  return int(response.Seguimiento)
}

func read_and_request_retail(conn *grpc.ClientConn,id int)(int){
  c := comms.NewCommsClient(conn)
  csvfilez ,_:= os.Open("../Pedidos/retail.csv")
  readerz := csv.NewReader(bufio.NewReader(csvfilez))
  for i:=0; true ;i++{
    line,error :=readerz.Read()
    if error==io.EOF{
      return -1
    }else if error!=nil{
        log.Fatal(error)
        continue
    }
    if i==id{
      aux1,_:=strconv.Atoi(line[2])
      response, err := c.CrearOrdenRetail(context.Background(),&comms.Request_CrearOrdenRetail{
        Id:line[0],
        Producto:line[1],
        Valor:int32(aux1),
        Tienda:line[3],
        Destino:line[4]})
      if err != nil {
        log.Fatalf("Error when calling SayHello: %s", err)
      }
      break
    }
  }
  csvfilez.Close()
  log.Printf("Numero de seguimento: %d", int(response.Seguimiento))
  return int(response.Seguimiento)
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
  conn, err := grpc.Dial("dist93:9000", grpc.WithInsecure())
  if err != nil {
    log.Fatalf("did not connect: %s", err)
  }
  defer conn.Close()
  var codigos [100]int
  var entregados int
  var opcion int
  var timex int
  var tipo int
  log.Printf("Bienvenido! Ingrese el numero de su tipo")
  log.Printf("1-Retail")
  log.Printf("2-Pyme")
  fmt.Scanln(&tipo)
  entregados=0
  log.Printf("Ingrese el tiempo entre inputs del cliente:")
  fmt.Scanln(&timex)
  var returnsis int
  for{
    time.Sleep(time.Duration(timex) * time.Second)
    opcion=rand.Intn(2)
    switch opcion {
    case 0:
      if(tipo=="1"){
        if(entregados<100){
          returnsis=read_and_request_retail(conn,entregados)
          if(returnsis!=-1){
            codigos[entregados]
            entregados++
          }
        }
      }else{
        returnsis=read_and_request_pymes(conn,entregados)
        if(returnsis!=-1){
          codigos[entregados]
          entregados++
        }
      }

    case 1:
      if(entregados>0){
        send_seguimento(conn,codigos[rand.Intn(entregados)])
      }
    }
  }
}

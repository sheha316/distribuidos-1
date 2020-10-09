package comms
import (
  "log"
  "golang.org/x/net/context"
  "os"
  "encoding/csv"
  "fmt"
  "strconv"
)

type Server struct {
}

type Pedido_pymes_l struct{
  Id string
  Producto string
  Valor int
  Tienda string
  Destino string
  Prioritario int
  Estado string
}

type Pedido_retail_l struct{
  Id string
  Producto string
  Valor int
  Tienda string
  Destino string
  Estado string
}
func (s *Server) CrearOrdenPyme(ctx context.Context, request *Request_CrearOrdenPyme) (*Response_CrearOrden, error) {
  log.Printf("Receive message %s", request.Id)

  seguimento:=0
  file,erros:=os.Open("./paquetes/1"+strconv.Itoa(seguimento)+".csv")
  for erros==nil{
    seguimento++
    file.Close()
    file,erros=os.Open("./paquetes/1"+strconv.Itoa(seguimento)+".csv")
  }
  file,erros=os.Create("./paquetes/1"+strconv.Itoa(seguimento)+".csv")
  if erros!=nil{
    log.Printf("error:")
    fmt.Println(erros)
    log.Printf("fin:")
  }
  writer:=csv.NewWriter(file)
  var guardar = [][]string{
    {request.Id,request.Producto,string(request.Valor),request.Tienda,request.Destino,string(request.Prioritario),"En bodega"},
  }
  erros=writer.WriteAll(guardar)
  aux:="1"+strconv.Itoa(seguimento)
  seguimento,_=strconv.Atoi(aux)
  file.Close()
  return &Response_CrearOrden{Seguimiento: int32(seguimento)}, nil
}

func (s *Server) CrearOrdenRetail(ctx context.Context, request *Request_CrearOrdenRetail) (*Response_CrearOrden, error) {
  log.Printf("Receive message %s", request.Id)

  seguimento:=0
  file,erros:=os.Open("./paquetes/2"+strconv.Itoa(seguimento)+".csv")
  for erros==nil{
    seguimento++
    file.Close()
    file,erros=os.Open("./paquetes/2"+strconv.Itoa(seguimento)+".csv")
  }
  file,erros=os.Create("./paquetes/2"+strconv.Itoa(seguimento)+".csv")
  if erros!=nil{
    log.Printf("error:")
    fmt.Println(erros)
    log.Printf("fin:")
  }
  writer:=csv.NewWriter(file)
  var guardar = [][]string{
    {request.Id,request.Producto,string(request.Valor),request.Tienda,request.Destino,"En bodega"},
  }
  erros=writer.WriteAll(guardar)
  aux:="2"+strconv.Itoa(seguimento)
  seguimento,_=strconv.Atoi(aux)
  file.Close()
  return &Response_CrearOrden{Seguimiento: int32(seguimento)}, nil
}

func (s *Server) Seguimiento(ctx context.Context, request *Request_Seguimiento) (*Response_Seguimiento, error) {
  log.Printf("Receive message %s", request.Seguimiento)
  aux:=strconv.Itoa(request.Seguimiento)
  csvFile,error:=os.Open("Prueba/"+aux+".csv")
  if error !=nil{
    return &Response_Seguimiento{Estado: "Paquete no existe"}, nil
  }
  reader := csv.NewReader(bufio.NewReader(csvFile))
  switch  aux[0]{
    case "1":
      var pedido []Pedido_pymes_l
    default:
      var pedido []Pedido_retail_l
  }
  for{
    line,error :=reader.Read()
    if error==io.EOF{
      break
    }else if error!=nil{
      log.Fatal(error)
    }
    aux1,_:=strconv.Atoi(line[2])
    switch  aux[0]{
      case "1":
        aux2,_:=strconv.Atoi(line[5])
        pedido=append(pedido,Pedido_retail{
          Id:line[0],
          Producto:line[1],
          Valor:aux1,
          Tienda:line[3],
          Destino:line[4],
          Prioritario:aux2,
          Estado:line[6],
        })
      default:
        pedido=append(pedido,Pedido_retail{
          Id:line[0],
          Producto:line[1],
          Valor:aux1,
          Tienda:line[3],
          Destino:line[4],
          Estado:line[5],
        })
    }
  }
  return &Response_Seguimiento{Estado: pedido[0].Estado}, nil
}

func (s *Server) SolicitarPaquete(ctx context.Context, request *Request_SolicitarPaquete) (*Response_SolicitarPaquete, error) {
  log.Printf("Receive message %s", request.Tipo)
  return &Response_SolicitarPaquete{Id: "vjdsv", Tipo: "nvjnvkjf", Valor: int32(7593), Tienda: "dsbjd", Destino: "bvjhdbvs"}, nil
}

func (s *Server) InformarEstado(ctx context.Context, request *Request_Estado) (*Response_Estado, error) {
  log.Printf("Receive message %s", request.Id)
  return &Response_Estado{Recibido: "holo"}, nil
}

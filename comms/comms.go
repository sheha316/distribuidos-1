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
  file,erros:=os.Open("./paquetes/0"+strconv.Itoa(seguimento)+".csv")
  for erros==nil{
    seguimento++
    file.Close()
    file,erros=os.Open("./paquetes/0"+strconv.Itoa(seguimento)+".csv")
  }
  file,erros=os.Create("./paquetes/0"+strconv.Itoa(seguimento)+".csv")
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
  aux:="0"+strconv.Itoa(seguimento)
  seguimento,_=strconv.Atoi(aux)
  file.Close()
  return &Response_CrearOrden{Seguimiento: int32(seguimento)}, nil
}

func (s *Server) Seguimiento(ctx context.Context, request *Request_Seguimiento) (*Response_Seguimiento, error) {
  log.Printf("Receive message %s", request.Seguimiento)
  return &Response_Seguimiento{Estado: "un estado muy bonito"}, nil
}

func (s *Server) SolicitarPaquete(ctx context.Context, request *Request_SolicitarPaquete) (*Response_SolicitarPaquete, error) {
  log.Printf("Receive message %s", request.Tipo)
  return &Response_SolicitarPaquete{Id: "vjdsv", Tipo: "nvjnvkjf", Valor: int32(7593), Tienda: "dsbjd", Destino: "bvjhdbvs"}, nil
}

func (s *Server) InformarEstado(ctx context.Context, request *Request_Estado) (*Response_Estado, error) {
  log.Printf("Receive message %s", request.Id)
  return &Response_Estado{Recibido: "holo"}, nil
}

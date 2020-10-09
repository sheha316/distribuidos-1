package comms
import (
  "log"
  "golang.org/x/net/context"
  "os"
  "encoding/csv"
)

type Server struct {
}

string id = 1;
string producto = 2;
int32 valor = 3;
string tienda = 4;
string destino = 5;
int32 prioritario = 6;

func (s *Server) CrearOrdenPyme(ctx context.Context, request *Request_CrearOrdenPyme) (*Response_CrearOrden, error) {
  log.Printf("Receive message %s", request.Id)

  seguimento:=0
  file,erros:=os.Open("./paquetes/"+string(seguimento)+".csv")
  while(erros==nil){
    seguimento++
    file,erros=os.Open("./paquetes/"+string(seguimento)+".csv")
  }
  file,erros:=os.Create("./paquetes/"+string(seguimento)+".csv")
  if erros!=nil{
    fmt.Println(erros)
  }
  writer:=csv.NewWriter(file)
  var guardar = [][]string{
    {request.Id,request.Producto,string(request.Valor),request.Tienda,request.Destino,string(request.Prioritario),}
  }
  erros=writer,WriteAll(guardar)
  return &Response_CrearOrden{Seguimiento: seguimiento}, nil
}

func (s *Server) CrearOrdenRetail(ctx context.Context, request *Request_CrearOrdenRetail) (*Response_CrearOrden, error) {
  log.Printf("Receive message %s", request.Id)
  seguimiento := int32(1)
  return &Response_CrearOrden{Seguimiento: seguimiento}, nil
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

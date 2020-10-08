package comms
import (
  "log"
  "golang.org/x/net/context"
)

type Server struct {
}


func (s *Server) CrearOrdenPyme(ctx context.Context, request *Request_CrearOrdenPyme) (*Response_CrearOrden, error) {
  log.Printf("Receive message %s", request.Id)
  seguimiento := int32(1)
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

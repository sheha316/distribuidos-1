package comms
import (
  "log"
  "golang.org/x/net/context"
)

type Server struct {
}

// string id = 1;
// string producto = 2;
// int valor = 3;
// string tienda = 4;
// string destino = 5;
// int prioritario = 6;

seguimiento int := 0

func (s *Server) CrearOrdenPyme(ctx context.Context, request *Request_CrearOrdenPyme) (*Response_CrearOrden, error) {
  log.Printf("Receive message %s", request.id)
  seguimiento += 1
  return &Response_CrearOrden{Seguimiento: "seguimiento"}, nil
}

func (s *Server) CrearOrdenRetail(ctx context.Context, request *Request_CrearOrdenRetail) (*Response_CrearOrden, error) {
  log.Printf("Receive message %s", request.id)
  seguimiento += 1
  return &Response_CrearOrden{Seguimiento: "seguimiento"}, nil
}

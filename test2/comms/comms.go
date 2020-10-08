package comms
import (
  "log"
  "golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, in *Request) (*Response, error) {
  log.Printf("Receive message %s", in.Greeting)
  return &Response{Greeting: "bar"}, nil
}

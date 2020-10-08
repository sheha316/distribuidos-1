package main
import (
  "log"
  "net"
  "github.com/sheha316/distribuidos-1/test2/comms"
  "google.golang.org/grpc"
)

func main() {
  lis, err := net.Listen("tcp", ":9000")
  if err != nil {
    log.Fatalf("failed to listen: %v", err)
  }

  s := comms.Server{}

  grpcServer := grpc.NewServer()

  comms.RegisterCommsServer(grpcServer, &s)
  if err := grpcServer.Serve(lis); err != nil {
    log.Fatalf("failed to serve: %s", err)
  }
}

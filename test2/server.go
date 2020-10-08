package main
import (
  "log"
  "net"
  "google.golang.org/grpc"
)

func main() {
  lis, err := net.Listen("tcp", ":9000")
  if err != nil {
    log.Fatalf("failed to listen: %v", err)
  }

  s := comms.Server{}

  grpcServer := grpc.NewServer()

  comms.RegisterPingServer(grpcServer, &s)
  if err := grpcServer.Serve(lis); err != nil {
    log.Fatalf("failed to serve: %s", err)
  }
}

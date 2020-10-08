package main
import (
  "log"
  "github.com/sheha316/distribuidos-1/comms"
  "golang.org/x/net/context"
  "google.golang.org/grpc"
)
func main() {
  var conn *grpc.ClientConn
  conn, err := grpc.Dial("dist93:9000", grpc.WithInsecure())
  if err != nil {
    log.Fatalf("did not connect: %s", err)
  }
  defer conn.Close()
  c := comms.NewCommsClient(conn)
  response, err := c.SayHello(context.Background(), &comms.Request{Greeting: "foo"})
  if err != nil {
    log.Fatalf("Error when calling SayHello: %s", err)
  }
  log.Printf("Response from server: %s", response.Estado)
}

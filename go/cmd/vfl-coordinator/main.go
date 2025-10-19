package main

import (
  "log"
  "net"

  "google.golang.org/grpc"
  v1 "github.com/Pedja-Djape/vfl-zero/go/pkg/protos/v1"
  "github.com/Pedja-Djape/vfl-zero/go/internal/api"
)

func main() {
  lis, err := net.Listen("tcp", ":8443")
  if err != nil { log.Fatalf("listen: %v", err) }

  s := grpc.NewServer() // TLS later
  v1.RegisterCoordinatorServer(s, api.NewCoordinatorServer())

  log.Println("coordinator listening on :8443")
  if err := s.Serve(lis); err != nil { log.Fatalf("serve: %v", err) }
}

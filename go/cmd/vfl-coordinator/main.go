package main

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"net"
	"os"

	"github.com/Pedja-Djape/vfl-zero/go/internal/api"
	v1 "github.com/Pedja-Djape/vfl-zero/go/pkg/protos/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
  lis, err := net.Listen("tcp", ":8443")
  if err != nil { log.Fatalf("listen: %v", err) }

  // load server cert and key
  cert, err := tls.LoadX509KeyPair("certs/server.crt", "certs/server.key");
  if err != nil { log.Fatalf("load key pair: %v", err) }

  caPEM, err := os.ReadFile("certs/dev-ca.crt")
  if err != nil { log.Fatalf("read CA cert: %v", err) }
  cp := x509.NewCertPool();
  if !cp.AppendCertsFromPEM(caPEM) { log.Fatalf("failed to append CA certs") }

  tlsCfg := &tls.Config{
    Certificates: []tls.Certificate{cert},
    MinVersion: tls.VersionTLS12,
    NextProtos: []string{"h2"},
    ClientCAs: cp,
    ClientAuth: tls.RequireAndVerifyClientCert,
  }

  s := grpc.NewServer(grpc.Creds(credentials.NewTLS(tlsCfg)))
  v1.RegisterCoordinatorServer(s, api.NewCoordinatorServer())

  log.Println("coordinator listening on :8443")
  if err := s.Serve(lis); err != nil { log.Fatalf("serve: %v", err) }
}

package api

import (
	"context"
	"fmt"

	v1 "github.com/Pedja-Djape/vfl-zero/go/pkg/protos/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/peer"
)

type coordinatorServer struct{ v1.UnimplementedCoordinatorServer }

func NewCoordinatorServer() *coordinatorServer { return &coordinatorServer{} }

func (s *coordinatorServer) RegisterParty(ctx context.Context, req *v1.RegisterPartyRequest) (*v1.RegisterPartyResponse, error) {
  p, ok := peer.FromContext(ctx);

  if !ok {
    return nil, status.Error(codes.Unauthenticated, "no peer found in context")
  }

  ti, ok := p.AuthInfo.(credentials.TLSInfo);
  if !ok || len(ti.State.PeerCertificates) == 0 {
    return nil, status.Error(codes.Unauthenticated, "no client certificate found")
  }
  cn := ti.State.PeerCertificates[0].Subject.CommonName;
  if cn != req.GetPartyId() {
    return nil, status.Error(codes.Unauthenticated, "client certificate does not match party ID")
  }

  return &v1.RegisterPartyResponse{
    SessionId: "sess_dev_0001",
    Echo:      fmt.Sprintf("hello, %s", req.GetPartyId()),
  }, nil
}

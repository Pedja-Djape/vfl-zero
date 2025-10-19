package api

import (
  "context"
  "fmt"

  v1 "github.com/Pedja-Djape/vfl-zero/go/pkg/protos/v1"
)

type coordinatorServer struct{ v1.UnimplementedCoordinatorServer }

func NewCoordinatorServer() *coordinatorServer { return &coordinatorServer{} }

func (s *coordinatorServer) RegisterParty(ctx context.Context, req *v1.RegisterPartyRequest) (*v1.RegisterPartyResponse, error) {
  return &v1.RegisterPartyResponse{
    SessionId: "sess_dev_0001",
    Echo:      fmt.Sprintf("hello, %s", req.GetPartyId()),
  }, nil
}

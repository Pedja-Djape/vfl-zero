package api

import (
	"context"
	"fmt"
	"sync"
	"time"

	v1 "github.com/Pedja-Djape/vfl-zero/go/pkg/protos/vfl/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

type coordinatorServer struct{ 
  v1.UnimplementedCoordinatorServer
  mu sync.Mutex
  parties map[string]string // party_id -> session_id
}

func NewCoordinatorServer() *coordinatorServer { return &coordinatorServer{ parties: make(map[string]string)} }

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
  s.mu.Lock()
  s.parties[cn] = "sess_dev_0001"
  s.mu.Unlock()

  return &v1.RegisterPartyResponse{
    SessionId: "sess_dev_0001",
    Echo:      fmt.Sprintf("hello, %s", req.GetPartyId()),
  }, nil
}

func (s *coordinatorServer) Heartbeat(ctx context.Context, req *v1.HeartbeatRequest) (*v1.HeartbeatResponse, error) {
  p, ok := peer.FromContext(ctx);
  if !ok {
    return nil, status.Error(codes.Unauthenticated, "no peer found in context")
  }

  ti, ok := p.AuthInfo.(credentials.TLSInfo);
  if !ok || len(ti.State.PeerCertificates) == 0 {
    return nil, status.Error(codes.Unauthenticated, "no client certificates found")
  }

  cn := ti.State.PeerCertificates[0].Subject.CommonName;

  s.mu.Lock()
  session, exists := s.parties[cn]
  s.mu.Unlock()

  if !exists {
    return nil, status.Error(codes.PermissionDenied, "party not registered")
  }

  if session != req.GetSessionId() {
    return nil, status.Error(codes.PermissionDenied, "session mismatch")
  }

  return &v1.HeartbeatResponse{
    Status: "ok",
    ServerUnixNano: time.Now().UTC().UnixNano(),
  }, nil
}

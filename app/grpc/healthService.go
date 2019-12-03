package grpc

import (
	"context"
	"sync"

	v1 "github.com/speix/memcash/app/pb/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Health struct {
	mu        sync.Mutex // protects statusMap
	statusMap map[string]v1.HealthCheckResponse_ServingStatus
}

func (s *Health) Check(ctx context.Context, req *v1.HealthCheckRequest) (*v1.HealthCheckResponse, error) {

	s.mu.Lock()
	defer s.mu.Unlock()

	if sts, ok := s.statusMap[req.Service]; ok {
		return &v1.HealthCheckResponse{
			Status: sts,
		}, nil
	}

	return nil, status.Error(codes.NotFound, "unknown service")
}

func NewHealthService(server *grpc.Server) *Health {

	health := &Health{
		statusMap: make(map[string]v1.HealthCheckResponse_ServingStatus),
	}

	health.statusMap[""] = v1.HealthCheckResponse_SERVING

	for s := range server.GetServiceInfo() {
		health.statusMap[s] = v1.HealthCheckResponse_SERVING
	}

	return health
}

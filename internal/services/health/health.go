package health_service

import (
	"context"

	"github.com/ErenDursun/grpc-server-template/api/grpc/health/v1"
)

type HealthServer struct {
	health.UnimplementedHealthServer
}

func (s *HealthServer) Check(ctx context.Context, in *health.HealthCheckRequest) (*health.HealthCheckResponse, error) {
	return &health.HealthCheckResponse{Status: health.HealthCheckResponse_UNKNOWN}, nil
}

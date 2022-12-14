package server

import (
	"context"

	"github.com/ErenDursun/grpc-server-template/api/grpc/health/v1"
)

type healthServer struct {
	health.UnimplementedHealthServer
}

func (s *healthServer) Check(ctx context.Context, in *health.HealthCheckRequest) (*health.HealthCheckResponse, error) {
	return &health.HealthCheckResponse{Status: health.HealthCheckResponse_UNKNOWN}, nil
}

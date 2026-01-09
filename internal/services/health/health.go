package health_service

import (
	"context"

	health "github.com/ErenDursun/grpc-server-template/api/gen/proto/health/v1"
)

type HealthServer struct {
	health.UnimplementedHealthServer
}

func (s *HealthServer) Check(ctx context.Context, in *health.HealthCheckRequest) (*health.HealthCheckResponse, error) {
	return &health.HealthCheckResponse{Status: health.HealthCheckResponse_UNKNOWN}, nil
}

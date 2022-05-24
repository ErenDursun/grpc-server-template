package server

import (
	"context"
	"log"

	"github.com/ErenDursun/grpc-server-template/api/grpc/echo/v1"
)

type echoServer struct {
	echo.UnimplementedEchoServer
}

func (s *echoServer) Echo(ctx context.Context, in *echo.EchoRequest) (*echo.EchoResponse, error) {
	log.Printf("Received: %v", in.Message)
	return &echo.EchoResponse{Message: in.Message}, nil
}

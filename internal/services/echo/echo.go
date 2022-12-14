package echo_service

import (
	"context"
	"fmt"

	"github.com/ErenDursun/grpc-server-template/api/grpc/echo/v1"
	"github.com/ErenDursun/grpc-server-template/internal/auth"
)

type EchoServer struct {
	echo.UnimplementedEchoServer
}

func (s *EchoServer) Echo(ctx context.Context, in *echo.EchoRequest) (*echo.EchoResponse, error) {
	if claims := auth.GetClaims(ctx); claims != nil {
		fmt.Printf("echo from authorized user '%v'\n", claims.Name)
	} else {
		fmt.Println("echo from unauthorized user")
	}

	return &echo.EchoResponse{Message: in.Message}, nil
}

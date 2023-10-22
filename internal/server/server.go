package server

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	grpc_jwt "github.com/ErenDursun/go-grpc-jwt-middleware/jwt"
	"github.com/golang-jwt/jwt/v5"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/ErenDursun/grpc-server-template/api/grpc/echo/v1"
	"github.com/ErenDursun/grpc-server-template/api/grpc/health/v1"
	"github.com/ErenDursun/grpc-server-template/internal/auth"
	echo_service "github.com/ErenDursun/grpc-server-template/internal/services/echo"
	health_service "github.com/ErenDursun/grpc-server-template/internal/services/health"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func NewServer() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	jwtConfig := grpc_jwt.Config{
		SigningKey: []byte("my_secret_key"),
		NewClaimsFunc: func(c context.Context) jwt.Claims {
			return &auth.GrpcServerClaims{}
		},
	}
	authFunc := grpc_jwt.NewAuthFuncWithConfig(jwtConfig)
	s := grpc.NewServer(
		grpc.StreamInterceptor(
			grpc_middleware.ChainStreamServer(
				grpc_ctxtags.StreamServerInterceptor(),
				grpc_auth.StreamServerInterceptor(authFunc),
				grpc_recovery.StreamServerInterceptor(),
			),
		),
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				grpc_ctxtags.UnaryServerInterceptor(),
				grpc_auth.UnaryServerInterceptor(authFunc),
				grpc_recovery.UnaryServerInterceptor(),
			),
		),
	)
	reflection.Register(s)
	health.RegisterHealthServer(s, &health_service.HealthServer{})
	echo.RegisterEchoServer(s, &echo_service.EchoServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

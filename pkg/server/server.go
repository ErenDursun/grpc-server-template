package server

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/ErenDursun/grpc-server-template/api/grpc/echo/v1"
	"github.com/ErenDursun/grpc-server-template/api/grpc/health/v1"
	"github.com/ErenDursun/grpc-server-template/internal/auth"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_jwt "github.com/grpc-ecosystem/go-grpc-middleware/auth/jwt"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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
	jwtConfig := grpc_jwt.JWTConfig{
		SigningKey: []byte("my_secret_key"),
		Claims:     &auth.GrpcServerClaims{},
	}
	authInterceptor := grpc_jwt.NewAuthFuncWithConfig(jwtConfig)
	s := grpc.NewServer(
		grpc.StreamInterceptor(
			grpc_middleware.ChainStreamServer(
				grpc_ctxtags.StreamServerInterceptor(),
				grpc_auth.StreamServerInterceptor(authInterceptor),
				grpc_recovery.StreamServerInterceptor(),
			),
		),
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				grpc_ctxtags.UnaryServerInterceptor(),
				grpc_auth.UnaryServerInterceptor(authInterceptor),
				grpc_recovery.UnaryServerInterceptor(),
			),
		),
	)
	reflection.Register(s)
	health.RegisterHealthServer(s, &healthServer{})
	echo.RegisterEchoServer(s, &echoServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

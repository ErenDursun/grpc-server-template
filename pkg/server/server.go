package server

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/ErenDursun/grpc-server-template/api/grpc/echo/v1"
	"github.com/ErenDursun/grpc-server-template/api/grpc/health/v1"
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
	s := grpc.NewServer()
	reflection.Register(s)
	health.RegisterHealthServer(s, &healthServer{})
	echo.RegisterEchoServer(s, &echoServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

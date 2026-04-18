package grpcinfra

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

type GRPCServer struct {
	server *grpc.Server
	port   string
}

func NewGRPCServer(port string) *GRPCServer {

	s := grpc.NewServer()

	return &GRPCServer{
		server: s,
		port:   port,
	}
}

func (g *GRPCServer) RegisterService(register func(*grpc.Server)) {
	register(g.server)
}

func (g *GRPCServer) Start() error {

	lis, err := net.Listen("tcp", ":"+g.port)
	if err != nil {
		return err
	}

	log.Println("gRPC server running on port", g.port)

	return g.server.Serve(lis)
}

func (g *GRPCServer) Stop(ctx context.Context) {
	log.Println("stopping gRPC server")
	g.server.GracefulStop()
}


package main

import (
	"context"
	"flag"
	"fmt"
	"net"

	pb "github.com/RaphaelPour/grpc-test/proto"
	"google.golang.org/grpc"
)

var (
	Port = flag.Int("port", 9000, "Server port")
)

type Server struct {
	pb.UnimplementedServerServer
}

func (s *Server) GetInfo(ctx context.Context, _ *pb.EmptyRequest) (*pb.ServerInfo, error) {
	return &pb.ServerInfo{
		Name:        "grpc-test",
		Author:      "evilc00kie",
		Description: "Test service to get server static information",
	}, nil
}

func main() {
	flag.Parse()
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *Port))
	if err != nil {
		fmt.Printf("error listening to localhost:%d: %s\n", *Port, err)
		return
	}

	grpcServer := grpc.NewServer()
	pb.RegisterServerServer(grpcServer, &Server{})
	grpcServer.Serve(listener)
}

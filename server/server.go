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

	Infos []pb.ServerInfo
}

func NewServer() *Server {
	server := &Server{
		Infos: make([]pb.ServerInfo, 0),
	}

	server.Infos = append(server.Infos, pb.ServerInfo{
		Name:        "grpc-test",
		Author:      "evilc00kie",
		Description: "Test service to get server static information",
	})

	return server
}

func (s *Server) GetInfo(ctx context.Context, _ *pb.EmptyRequest) (*pb.ServerInfo, error) {
	return &pb.ServerInfo{
		Name:        "grpc-test",
		Author:      "evilc00kie",
		Description: "Test service to get server static information",
	}, nil
}

func (s *Server) AddServer(ctx context.Context, info *pb.ServerInfo) (*pb.EmptyResponse, error) {
	s.Infos = append(s.Infos, *info)
	fmt.Printf("[*] New server added: %s\n", info.Name)
	return &pb.EmptyResponse{}, nil
}

func main() {
	flag.Parse()
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *Port))
	if err != nil {
		fmt.Printf("error listening to localhost:%d: %s\n", *Port, err)
		return
	}

	grpcServer := grpc.NewServer()
	pb.RegisterServerServer(grpcServer, NewServer())
	grpcServer.Serve(listener)
}

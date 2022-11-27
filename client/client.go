package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	pb "github.com/RaphaelPour/grpc-test/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	Port = flag.Int("port", 9000, "Server port")
)

func main() {
	flag.Parse()

	connection, err := grpc.Dial(fmt.Sprintf("localhost:%d", *Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("error connecting to localhost:%d: %s\n", *Port, err)
		return
	}
	defer connection.Close()
	client := pb.NewServerClient(connection)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	serverInfo, err := client.GetInfo(ctx, &pb.EmptyRequest{})
	if err != nil {
		fmt.Printf("error getting server info: %s\n", err)
		return
	}
	fmt.Println("Name:", serverInfo.Name)
	fmt.Println("Description:", serverInfo.Description)
	fmt.Println("Author:", serverInfo.Author)
}

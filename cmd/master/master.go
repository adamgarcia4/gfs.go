package main

import (
	"context"
	"log"
	"net"

	pb "github.com/adamgarcia4/gfs.go/api/proto"
	"google.golang.org/grpc"
)

/**

Responsible for orchestrating all chunk info
*/

// The first thing I need to do is receive a CreateFileRequest
type masterServer struct {
	pb.UnimplementedChatServiceServer
}

// GetFeature(ctx context.Context, point *pb.Point) (*pb.Feature, error) {
// rpc SayHello(Message) returns (Message) {}
func (m masterServer) SayHello(ctx context.Context, message *pb.Message) (*pb.Message, error) {
	return &pb.Message{
		Body: "Hello",
	}, nil
}

func main() {
	// print out the message
	log.Println("Starting gRPC server on port 9000...")

	// check that masterServer implements the pb.ChatServiceServer interface
	var _ pb.ChatServiceServer = &masterServer{}

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterChatServiceServer(grpcServer, &masterServer{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9000: %v", err)
	}
}

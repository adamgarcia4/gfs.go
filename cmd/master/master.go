package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

/**

Responsible for orchestrating all chunk info
*/

// The first thing I need to do is receive a CreateFileRequest

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	grpcServer := grpc.NewServer()

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9000: %v", err)
	}
}

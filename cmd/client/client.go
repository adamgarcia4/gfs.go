package main

import (
	"context"
	"log"

	pb "github.com/adamgarcia4/gfs.go/api/proto"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	grpc.NewClient(":9000", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}
	defer conn.Close()

	c := pb.ChatServiceClient(conn)

	message := &pb.Message{
		Body: "Hello from the client!",
	}

	response, err := c.SayHello(context.Background(), message)
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}

	log.Printf("Response from the server: %s", response.Body)
}

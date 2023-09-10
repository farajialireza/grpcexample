package main

import (
	pb "grpcexample/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewMyServiceClient(conn)
	name := "John"
	req := &pb.HelloRequest{Name: name}
	resp, err := client.SayHello(context.Background(), req)
	if err != nil {
		log.Fatalf("Error calling SayHello: %v", err)
	}
	fmt.Printf("Response from server: %s\n", resp.GetMessage())
}

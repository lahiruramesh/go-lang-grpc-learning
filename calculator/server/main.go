package main

import (
	"log"
	"net"
	"google.golang.org/grpc"
	pb "github.com/lahiruramesh/go-grpc-learn/calculator/proto"
)

var addr string = "0.0.0.0:50051"

type server struct {
	pb.CalculatorServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)

	// Create a new gRPC server
	s := grpc.NewServer()

	pb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
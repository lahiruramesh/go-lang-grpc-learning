package main

import (
	"context"
	"log"

	pb "github.com/lahiruramesh/go-grpc-learn/greet/proto"
)

func (s *server) Greet(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Received: %v\n", req)

	return &pb.GreetResponse{Result: "Hello" + req.FirstName}, nil
}
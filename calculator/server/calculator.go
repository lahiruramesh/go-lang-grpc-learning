package main

import (
	"context"
	"log"

	pb "github.com/lahiruramesh/go-grpc-learn/calculator/proto"
)

func (s *server) Add(ctx context.Context, req *pb.CalculatorRequest) (*pb.CalculatorResponse, error) {
	log.Printf("Received: %v\n", req)

	return &pb.CalculatorResponse{Result: req.A + req.B}, nil
}

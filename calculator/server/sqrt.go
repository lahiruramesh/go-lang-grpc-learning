package main

import (
	"context"
	"fmt"
	pb "github.com/lahiruramesh/go-grpc-learn/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"math"
)

func (s *server) Sqrt(ctx context.Context, req *pb.SqrtRequest) (*pb.SqrtResponse, error) {
	log.Printf("Received Sqrt RPC: %v\n", req)
	number := req.GetNumber()
	if number < 0 {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Received a negative number: %v", number))
	}
	return &pb.SqrtResponse{Result: math.Sqrt(float64(number))}, nil
}

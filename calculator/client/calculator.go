package main

import (
	"context"
	"log"

	pb "github.com/lahiruramesh/go-grpc-learn/calculator/proto"
)

func doCalc(c pb.CalculatorServiceClient) {
	req := &pb.CalculatorRequest{
		A: 10,
		B: 20,
	}
	ctx := context.Background()
	res, err := c.Add(ctx, req)

	if err != nil {
		log.Fatalf("Failed to add: %v\n", err)
	}

	log.Printf("Add response: %d\n", res.Result)
}
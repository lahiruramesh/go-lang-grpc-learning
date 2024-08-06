package main

import (
	pb "github.com/lahiruramesh/go-grpc-learn/greet/proto"
	"context"
	"log"
)

func doGreet(c pb.GreetServiceClient) {
	req := &pb.GreetRequest{
		FirstName: "Lahiru",
	}
	ctx := context.Background()
	res, err := c.Greet(ctx, req)

	if err != nil {
		log.Fatalf("Failed to greet: %v\n", err)
	}

	log.Printf("Greet response: %s\n", res.Result)
}
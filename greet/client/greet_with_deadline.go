package main

import (
	"context"
	"fmt"
	"log"
	"time"
	pb "github.com/lahiruramesh/go-grpc-learn/greet/proto"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
)

func doGreetWithDeadline(c pb.GreetServiceClient, timeout time.Duration) {
	log.Println("Starting to do a UnaryWithDeadline RPC...")
	req := &pb.GreetRequest{
		FirstName: "Lahiru",
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	res, err := c.GreetWithDeadline(ctx, req)
	if err != nil {
		status, ok := status.FromError(err)
		if ok {
			if status.Code() == codes.DeadlineExceeded {
				fmt.Println("Timeout was hit! Deadline was exceeded")
			} else {
				fmt.Printf("Unexpected error: %v", status)
			}
		} else {
			log.Fatalf("error while calling GreetWithDeadline RPC: %v", err)
		}
		return
	}

	log.Printf("Response from GreetWithDeadline: %v", res.Result)
}
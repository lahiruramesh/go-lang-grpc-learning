package main

import (
	"context"
	"log"
	"time"

	pb "github.com/lahiruramesh/go-grpc-learn/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) GreetWithDeadline(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Println("Received GreetWithDeadline RPC")

	for i := 0; i < 3; i++ {
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("The client canceled the request!")
			return nil, status.Error(codes.Canceled, "The client canceled the request!")
		}

		log.Printf("Sleep for 1 second...")
		time.Sleep(1 * time.Second)
	}

	firstName := req.GetFirstName()
	result := "Hello " + firstName + "! "

	return &pb.GreetResponse{
		Result: result,
	}, nil
}

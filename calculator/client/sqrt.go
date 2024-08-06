package main 

import (
	pb "github.com/lahiruramesh/go-grpc-learn/calculator/proto"
	"log"
	"context"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
)

func doSqrt(c pb.CalculatorServiceClient, number int32) {
	req := &pb.SqrtRequest{Number: number}
	res, err := c.Sqrt(context.Background(), req)
	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			log.Fatalf("Error while calling Sqrt RPC: %v", e.Message())
			log.Fatalf("Error code: %v", e.Code())

			if e.Code() == codes.InvalidArgument {
				log.Fatalf("We probably sent a negative number!")
				return
			}
		} else {
			log.Fatalf("Error while calling Sqrt RPC not grpc: %v", err)
		}
		
	}
	log.Printf("Response from Sqrt: %v", res.Result)
	
}
package main

import (
	"io"
	"log"
	"context"
	pb "github.com/lahiruramesh/go-grpc-learn/greet/proto"
)

func doGreetManyTimes(c pb.GreetServiceClient) {
	log.Println("Starting to do a GreetManyTimes RPC...")

	req := &pb.GreetRequest{
		FirstName: "Lahiru",
	}

	ctx := context.Background()

	stream, err := c.GreetManyTimes(ctx, req)

	if err != nil {
		log.Fatalf("Error while calling GreetManyTimes RPC: %v", err)
	}

	for {
		msg, err := stream.Recv()


		if err == io.EOF {
			log.Fatalf("Stream end: %v", err)
			break
		}

		if (err != nil) {
			log.Fatalf("Error while reading stream: %v", err)
		}

		log.Printf("Response from GreetManyTimes: %s/n", msg.GetResult())
	}
}
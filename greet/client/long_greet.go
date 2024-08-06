package main

import (
	"context"
	"log"
	"time"

	pb "github.com/lahiruramesh/go-grpc-learn/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("Starting to do a LongGreet RPC...");

	reqs := []*pb.GreetRequest{{FirstName: "Lahiru"}, {FirstName: "Ramesh"}, {FirstName: "Sanjeewa"}}

	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("Error while calling LongGreet RPC: %v", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response from LongGreet: %v", err)
	}

	log.Printf("LongGreet response: %v\n", res.Result)
}
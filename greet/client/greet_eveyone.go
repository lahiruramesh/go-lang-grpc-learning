package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/lahiruramesh/go-grpc-learn/greet/proto"
)

func doGreetEveryone(c pb.GreetServiceClient) {
	log.Println("Starting to do a GreetEveryone RPC...")

	stream, err := c.GreetEveryone(context.Background())

	if err != nil {
		log.Fatalf("Error while calling GreetEveryone RPC: %v", err)
	}

	reqs := []*pb.GreetRequest{{FirstName: "Lahiru"}, {FirstName: "Ramesh"}, {FirstName: "Sanjeewa"}}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Sending message: %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Error while receiving data: %v", err)
				break
			}

			log.Printf("Received: %v\n", res.GetResult())
		}
		close(waitc)
	}()

	<-waitc
}
package main

import (
	pb "github.com/lahiruramesh/go-grpc-learn/greet/proto"
	"io"
	"log"
)

func (s *server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Println("Received LongGreet RPC")
	result := ""
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: result,
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}

		log.Printf("Receiving %v!\n", req.FirstName)

		result += "Hello " + req.FirstName + "! "
	}
}

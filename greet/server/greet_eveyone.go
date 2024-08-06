package main 

import (
	"io"
	"log"
	pb "github.com/lahiruramesh/go-grpc-learn/greet/proto"
)

func (s *server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {
	log.Println("Received GreetEveryone RPC")
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
			return err
		}

		firstName := req.GetFirstName()
		result := "Hello " + firstName + "! "

		sendErr := stream.Send(&pb.GreetResponse{
			Result: result,
		})

		if sendErr != nil {
			log.Fatalf("Error while sending data to client: %v", sendErr)
		}
	}
}
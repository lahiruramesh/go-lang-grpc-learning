package main

import (
	pb "github.com/lahiruramesh/go-grpc-learn/greet/proto"
	"log"
)

func (s *server) GreetManyTimes(req *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	log.Printf("Received GreetManyTimes RPC: %v\n", req)
	firstName := req.FirstName

	for i := 0; i < 10; i++ {
		res := &pb.GreetResponse{
			Result: "Hello " + firstName + " number " + string(i),
		}
		stream.Send(res)
	}

	return nil
}
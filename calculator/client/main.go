package main

import (
	"log"
	"google.golang.org/grpc"
	pb "github.com/lahiruramesh/go-grpc-learn/calculator/proto"
)

var addr string = "localhost:50051";

func main() {

	con, err := grpc.Dial(addr, grpc.WithInsecure());

	if err != nil {
		log.Fatalf("Failed to dial server: %v\n", err);
	}

	defer con.Close();

	c := pb.NewCalculatorServiceClient(con);

	// doCalc(c);

	doSqrt(c, -10);

}
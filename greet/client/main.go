package main

import (
	"log"
	"time"

	pb "github.com/lahiruramesh/go-grpc-learn/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var addr string = "localhost:50051";

func main() {

	tls := true;
	opts := []grpc.DialOption{};

	if tls {
		certFile := "ssl/ca.crt";
		creds, sslErr := credentials.NewClientTLSFromFile(certFile, "");

		if sslErr != nil {
			log.Fatalf("Failed loading certificates: %v\n", sslErr);
			return;
		}

		opts = append(opts, grpc.WithTransportCredentials(creds));
	}

	conn, err := grpc.NewClient(addr, opts...);

	if err != nil {
		log.Fatalf("Failed to dial server: %v\n", err);
	}

	defer conn.Close();


	c := pb.NewGreetServiceClient(conn);

	// doGreet(c);

	//doGreetManyTimes(c);

	// doLongGreet(c);
	// doGreetEveryone(c);
	doGreetWithDeadline(c, 1*time.Second);

}
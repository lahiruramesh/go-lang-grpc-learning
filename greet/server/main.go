package main

import (
	"log"
	"net"
	pb "github.com/lahiruramesh/go-grpc-learn/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type server struct {
	pb.GreetServiceServer
}


var addr string = "0.0.0.0:50051"



func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)

	// Set up TLS
	tls := true
	opts := []grpc.ServerOption{}

	if tls {
		// Create a new gRPC server with TLS credentials
		// s := grpc.NewServer(grpc.Creds(credentials.NewServerTLSFromCert(&insecure.Cert)))
		// pb.RegisterGreetServiceServer(s, &server{})
		// if err := s.Serve(lis); err != nil {
		// 	log.Fatalf("Failed to serve: %v\n", err)
		// }

		certFile := "ssl/server.crt"
		keyFile := "ssl/server.pem"
		creds, sslErr := credentials.NewServerTLSFromFile(certFile, keyFile)

		if sslErr != nil {
			log.Fatalf("Failed loading certificates: %v", sslErr)
			return
		}

		opts = append(opts, grpc.Creds(creds))
	}
	// Create a new gRPC server
	s := grpc.NewServer(opts...)
	pb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}

}
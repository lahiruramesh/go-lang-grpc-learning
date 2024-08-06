package main

import (
	"context"
	"log"

	pb "github.com/lahiruramesh/go-grpc-learn/blog/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	client := pb.NewBlogServiceClient(conn)

	req := &pb.Blog{
		AuthorId: "Lahiru Ramesh",
		Title:    "My first blog",
		Content:  "This is my first blog. I am writing this to learn gRPC.",
	}

	res, err := client.CreateBlog(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to create blog: %v", err)
	}

	log.Printf("Blog created with id: %v", res.Id)		

}
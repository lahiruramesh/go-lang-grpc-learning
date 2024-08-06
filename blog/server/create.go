package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/lahiruramesh/go-grpc-learn/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *Server) CreateBlog(ctx context.Context, req *pb.Blog) (*pb.BlogId, error) {
	log.Printf("Create blog request received: %v", req)

	data := BlogItem{
		AuthorId: req.GetAuthorId(),
		Title:    req.GetTitle(),
		Content:  req.GetContent(),
	}

	res, err := collection.InsertOne(context.Background(), data)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)

	if !ok {
		return nil, status.Errorf(
			codes.Internal,
			"Cannot convert to OID",
		)
	}

	return &pb.BlogId{
		Id: oid.Hex(),
	}, nil
}

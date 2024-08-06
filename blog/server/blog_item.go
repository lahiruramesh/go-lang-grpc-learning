package main

import (
	pb "github.com/lahiruramesh/go-grpc-learn/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogItem struct {
	ID  primitive.ObjectID `bson:"_id,omitempty"`
	AuthorId string `bson:"author_id"`
	Title string `bson:"title"`
	Content string `bson:"content"`
}

func documentToBlogItem(doc *pb.Blog) *pb.Blog {
	return &pb.Blog{
		Id: doc.Id,
		AuthorId: doc.AuthorId,
		Title: doc.Title,
		Content: doc.Content,
	}
}
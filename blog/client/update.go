package main

import (
	"context"
	"log"

	pb "github.com/emreaslan/grpc-go/blog/proto"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Println("UpdateBlog was invoked")

	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "Aslan",
		Title:    "Updated title",
		Content:  "Updated content",
	}

	_, err := c.UpdateBlog(context.Background(), newBlog)

	if err != nil {
		log.Fatalf("Error while updating blog: %v\n", err)
		return
	}

	log.Println("Blog was updated!")
}

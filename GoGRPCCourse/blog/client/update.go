package main

import (
	"context"
	"log"

	pb "github.com/BugzTheBunny/GoCourses/blog/proto"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Println("---updateBlog was Invoked---")

	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "Not Bugz",
		Title:    "New Title",
		Content:  "Whole new content",
	}

	_, err := c.UpdateBlog(context.Background(), newBlog)

	if err != nil {
		log.Fatalf("Error happened while updating: %v\n", err)
	}

	log.Println("Log was updated")
}

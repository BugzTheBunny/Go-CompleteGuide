package main

import (
	"context"
	"log"

	pb "github.com/BugzTheBunny/GoCourses/blog/proto"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("---createBlog was Invoked---")

	blog := &pb.Blog{
		AuthorId: "Bugz",
		Title:    "First Blog",
		Content:  "Content for the first blog",
	}

	res, err := c.CreateBlog(context.Background(), blog)

	if err != nil {
		log.Fatalf("Unexpected error while creating blog: %v\n", err)
	}

	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}

	log.Printf("Blog has been created: %s\n", res)

	return res.Id
}

package main

import (
	"log"

	pb "github.com/BugzTheBunny/GoCourses/blog/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect %v", err)
	}

	defer conn.Close()

	c := pb.NewBlogServiceClient(conn)

	validId := createBlog(c)
	// invalidId := "invalid"
	readBlog(c, validId)
	// readBlog(c, invalidId)
	updateBlog(c, validId)
	listBlog(c)
	deleteBlog(c, validId)
}

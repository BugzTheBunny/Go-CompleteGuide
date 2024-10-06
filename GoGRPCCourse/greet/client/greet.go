package main

import (
	"context"
	"log"

	pb "github.com/BugzTheBunny/GoCourses/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("Do greet invoked")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Slava",
	})

	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}

	log.Printf("Greeting: %s\n", res.Result)
}

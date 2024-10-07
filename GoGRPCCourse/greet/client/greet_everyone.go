package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/BugzTheBunny/GoCourses/greet/proto"
)

func doGreetEveryone(c pb.GreetServiceClient) {
	log.Println("doGreetEveryone was Invoked")

	stream, err := c.GreetEveryone(context.Background())

	if err != nil {
		log.Fatalf("Error while creating stream: %v", err)
	}

	reqs := []*pb.GreetRequest{
		{FirstName: "Bugz"},
		{FirstName: "The"},
		{FirstName: "Bunny"},
	}

	waitc := make(chan chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Seding request: %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("error while receiving: %v\n", err)
				break
			}

			log.Printf("Received: %v", res.Result)
		}
		close(waitc)
	}()
	<-waitc
}

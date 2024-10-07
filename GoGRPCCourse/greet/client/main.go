package main

import (
	"log"

	pb "github.com/BugzTheBunny/GoCourses/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var addr string = "localhost:50051"

func main() {
	tls := true // Change to false if needed

	opts := []grpc.DialOption{}

	if tls {
		certFile := "ssl/ca.crt"
		creds, err := credentials.NewClientTLSFromFile(certFile, "")

		if err != nil {
			log.Fatalf("Error while loading CA trust certification: %v\n", err)
		}

		opts = append(opts, grpc.WithTransportCredentials(creds))
	}
	conn, err := grpc.NewClient(addr, opts...)

	if err != nil {
		log.Fatalf("Failed to connect %v", err)
	}

	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)

	doGreet(c)
	// doGreetManyTimes(c)
	// doLongGreet(c)
	// doGreetEveryone(c)
	// doGreetWithDeadline(c, 5*time.Second)
	// doGreetWithDeadline(c, 2*time.Second) // This wil fail

}

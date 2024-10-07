package main

import (
	"log"
	"net"

	pb "github.com/BugzTheBunny/GoCourses/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.UnimplementedGreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on %v", err)
	}

	log.Printf("Listening on %s\n", addr)

	opts := []grpc.ServerOption{}
	tls := true // Change to false if needed

	if tls {
		certFile := "ssl/server.crt"
		keyFile := "ssl/server.pem"
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)

		if err != nil {
			log.Fatalf("Failed to load certificates: %v\n", err)
		}

		opts = append(opts, grpc.Creds(creds))
	}

	s := grpc.NewServer(opts...)

	pb.RegisterGreetServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
}

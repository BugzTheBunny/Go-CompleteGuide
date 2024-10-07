package main

import (
	"log"

	pb "github.com/BugzTheBunny/GoCourses/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect %v", err)
	}

	defer conn.Close()

	c := pb.NewCalculatorServiceClient(conn)

	//doSum(c)
	// doPrimes(c)
	// doAvg(c)
	// doMax(c)

	doSqrt(c, 10)
	doSqrt(c, 15)
	doSqrt(c, -1) // Will fail

}

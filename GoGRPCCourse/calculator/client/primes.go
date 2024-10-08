package main

import (
	"context"
	"io"
	"log"

	pb "github.com/BugzTheBunny/GoCourses/calculator/proto"
)

func doPrimes(c pb.CalculatorServiceClient) {
	log.Println("doPrimes was Invoked")

	req := &pb.PrimeRequest{
		Number: 12390392840,
	}

	stream, err := c.Primes(context.Background(), req)

	if err != nil {
		log.Fatalf("error while calling Primes: %v\n", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream: %v", err)
		}

		log.Printf("Primes: %d\n", res.Result)
	}

}

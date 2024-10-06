package main

import (
	"fmt"
	"time"
)

func greet(phrase string, isDone chan bool) {
	fmt.Println("Hello! ", phrase)
	isDone <- true
}

func slowGreet(phrase string, isDone chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hello! ", phrase)
	isDone <- true
	close(isDone)
}

func main() {
	isDone := make(chan bool)
	go greet("Test 1", isDone)
	go greet("Test 2", isDone)
	go slowGreet("Test 3", isDone)
	go greet("Test 4", isDone)

	for range isDone {
	}
}

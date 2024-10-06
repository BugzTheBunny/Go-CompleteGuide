# Concurrency in Go

### - `Go`

You can execute the code at the same time, by adding `go` to it, like in the code below.  
Note that probably you will not see anything in the terminal, as when the main stops execution  
it will kill the rest of the executions, so we will never see the `slowGreet`, as it takes 3 seconds to execute.  

    ```
    func greet(phrase string) {
        fmt.Println("Hello! ", phrase)
    }

    func slowGreet(phrase string) {
        time.Sleep(3 * time.Second)
        fmt.Println("Hello! ", phrase)
    }
    func main() {
        go greet("Test 1")
        go greet("Test 2")
        go slowGreet("Test 3")
        go greet("Test 4")
    }
    ```

### - `Channels`

Channels are like "signals", in the code below, the channel `chan isDone` is sent to the function `slowGreet`, it will make the main function, to not finish execution, before we have the response.  

In big picture, its a communication device between functions.

```
package main

import (
	"fmt"
	"time"
)

func greet(phrase string) {
	fmt.Println("Hello! ", phrase)
}

func slowGreet(phrase string, doneChan chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hello! ", phrase)
	doneChan <- 5
}

func main() {
	go greet("Test 1")
	go greet("Test 2")
	isDone := make(chan int)
	go slowGreet("Test 3", isDone)
	go greet("Test 4")
	fmt.Println(<-isDone)
}

```

You can create an array of channels, and connecto to all of them, in this way, you will be able to use channels and goroutines, and at the same time follow up.

```
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
}

func main() {
	conditions := make([]chan bool, 4)
	conditions[0] = make(chan bool)
	go greet("Test 1", conditions[0])
	conditions[1] = make(chan bool)
	go greet("Test 2", conditions[1])
	conditions[2] = make(chan bool)
	go slowGreet("Test 3", conditions[2])
	conditions[3] = make(chan bool)
	go greet("Test 4", conditions[3])

	for _, done := range conditions {
		<-done
	}
}
```
- But there is a better way:  
    in the function below, we send all of the functions the same channel, and we loop over it, if everything there is done, and in the last function, meaning the one that will finish the last
    we `close` the channel - in our case, inside  `slowGreet`.  
    That way go knows that the everything that is working on this channel is done.

    ```
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
    ```
package main

import "fmt"

func main() {
	fact := factorial(3)
	fmt.Println(fact)
}

func factorial(number int) int {
	fmt.Println(number)
	if number == 0 {
		return 1
	}
	return number * factorial(number-1)
}

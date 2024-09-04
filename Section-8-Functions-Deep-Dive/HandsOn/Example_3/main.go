package main

import "fmt"

func main() {
	sum := sumup(1, 10, 20, 30, 40)
	sumDiff := sumDifferent("test", 12, 3, 4, 5, 6)
	fmt.Println(sum, sumDiff)
}

func sumup(numbers ...int) int {
	sum := 0
	for _, val := range numbers {
		sum += val
	}
	return sum
}

func sumDifferent(word string, numbers ...int) int {
	fmt.Println(word)
	fmt.Println(numbers)
	return 0
}

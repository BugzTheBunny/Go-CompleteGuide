package main

import "fmt"

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	f := func(number int) int { return number * 2 }
	numbersDoubled := transformNumbers(&numbers, f)
	fmt.Println(numbersDoubled)
}

func transformNumbers(numbers *[]int, transformFunc func(int) int) []int {
	dNumbers := []int{}
	for _, value := range *numbers {
		dNumbers = append(dNumbers, transformFunc(value))
	}
	return dNumbers
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}

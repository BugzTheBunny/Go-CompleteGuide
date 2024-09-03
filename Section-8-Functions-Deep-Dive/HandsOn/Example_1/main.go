package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4}
	numbersDoubled := transformNumbers(&numbers, double)
	fmt.Println(numbersDoubled)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}
	for _, value := range *numbers {
		dNumbers = append(dNumbers, transform(value))
	}
	return dNumbers
}

func double(number int) int {
	return number * 2
}

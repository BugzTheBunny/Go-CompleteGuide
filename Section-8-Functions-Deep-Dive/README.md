# Functions deep dive & Recursion
You can pass functions as variables, like in the example below:

```
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
```

But as you see, its not the best practice, because if a function should have a lot of values or returns, it can get complicated, therefore we can createe a type of a function.

```
package main

import "fmt"

type transformFn func(int) int

type anotherTransformFn func(int, string) int

func main() {
	numbers := []int{1, 2, 3, 4}
	numbersDoubled := transformNumbers(&numbers, double)
	fmt.Println(numbersDoubled)
}

func transformNumbers(numbers *[]int, transform transformFn) []int {
	dNumbers := []int{}
	for _, value := range *numbers {
		dNumbers = append(dNumbers, transform(value))
	}
	return dNumbers
}

func double(number int) int {
	return number * 2
}
```
You can also return a function as a value, for example, here we get `double` if the number is under 5 else we get `triple`
```
package main

import "fmt"

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	numbersDoubled := transformNumbers(&numbers)
	fmt.Println(numbersDoubled)
}

func transformNumbers(numbers *[]int) []int {
	dNumbers := []int{}
	for _, value := range *numbers {
		transformFunction := getTransformerFunction(5)
		dNumbers = append(dNumbers, transformFunction(value))
	}
	return dNumbers
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}

func getTransformerFunction(value int) transformFn {
	if value > 5 {
		return double
	} else {
		return triple
	}
}

```
There are also Anonymous functions
```
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

```
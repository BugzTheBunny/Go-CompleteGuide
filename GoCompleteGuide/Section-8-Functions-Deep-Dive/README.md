# Functions deep dive & Recursion
- You can pass functions as variables, like in the example below:

    ```
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

- But as you see, its not the best practice, because if a function should have a lot of values or returns, it can get complicated, therefore we can createe a type of a function.

    ```
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

- You can also return a function as a value, for example, here we get `double` if the number is under 5 else we get `triple`

    ```
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
    
- There are also Anonymous functions
    ```
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
    ```

- `Variadic Functions`
    If you dont know the exact number of arguments you will use, you can use variadic functions, like in the example below.
    ```
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

    ```
package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	courseRatings := make(floatMap, 3)
	courseRatings.output()

	users := make([]string, 2, 5)
	users[0] = "Tom"
	users = append(users, "Josh")
	users = append(users, "Tim")

	for index, value := range users {
		fmt.Println(index, " ", value)
	}
}

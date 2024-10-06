package main

import "fmt"

func main() {
	hobbies := [3]string{"Sports", "Books", "Board Games"}
	fmt.Println(hobbies)

	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:3])

	selectedHobbies := hobbies[:2]
	fmt.Println(selectedHobbies)

	selectedHobbies = selectedHobbies[1:3]
	fmt.Println(selectedHobbies)

	courseGoals := []string{"Do", "Something"}
	fmt.Println(courseGoals)

	courseGoals[1] = "Something cool"
	courseGoals = append(courseGoals, "Like very very cool")
	fmt.Println(courseGoals)
}

package main

import (
	"fmt"

	"example.com/app/user"
)

func main() {

	firstName := "Bugz"
	lastName := "Bunny"
	birthdate := "05/05/1555"

	appUser, err := user.NewUser(firstName, lastName, birthdate)

	if err != nil {
		fmt.Println("Something went wrong...")
		return
	}

	appUser.OutputUserDetails()
	appUser.ClearUserData()
	appUser.OutputUserDetails()
}

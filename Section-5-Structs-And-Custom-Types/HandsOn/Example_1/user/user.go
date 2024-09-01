package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

type Admin struct {
	User
	email    string
	password string
}

func (u User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthdate)

}

func (u *User) ClearUserData() {
	u.firstName = ""
	u.lastName = ""
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email: email, password: password, User: User{
			firstName: "Dummy",
			lastName:  "Dummy",
			birthdate: "----",
			createdAt: time.Now(),
		},
	}
}

func NewUser(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("fields are missing")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}

# Structs

Structs are essentialy the "classes" of Go.

# Struct example

```
package main

import (
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

func main() {
	firstName := "Bugz"
	lastName := "Bunny"
	birthdate := "05/05/1555"

	appUser := User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}
}
```

## Methods

    - Methods are **attached** to structs, for example like this:  
        `func (u User) outputUserDetails() `
    - Methods that should change the data in the struct, should be passed as pointers:  
        `func (u *User) clearUserData()`

```
package main

import (
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

func (u User) outputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthdate)

}

func (u *User) clearUserData() {
	u.firstName = ""
	u.lastName = ""
}

func main() {
	firstName := "Bugz"
	lastName := "Bunny"
	birthdate := "05/05/1555"

	appUser := User{
        ....Data...
	}

	appUser.outputUserDetails()
}
```

## Constructor

Constructors look like this:
```
func newUser(firstName, lastName, birthdate string) User {
	return User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}
}
```

But usually you will use it with pointers:
```
func newUser(firstName, lastName, birthdate string) *User {
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}
}
```

## Validation
When creating a new structure, in a constructor you can validate the the fields, for example:


**Constructor with validation:**
```
func newUser(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("Fields are missing")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}
```

## Composition
In order to use inheritance, you can just assign annonymously User for a new struct, like Admin, then it will inherit everything from User.
```
type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	User <----
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
```
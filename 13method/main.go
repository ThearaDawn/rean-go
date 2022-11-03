package main

import (
	"fmt"
)

func main() {

	fmt.Println("Let's go with Method in Golang")

	thereaD := User{"Theara", "theara@gmail.com", true, 18}

	fmt.Println(thereaD)
	fmt.Printf("User Info: %+v\n", thereaD)

	thereaD.userEmail()
	thereaD.updateUserEmail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (user User) userEmail() {
	fmt.Println("User email:", user.Email)
}

func (user User) updateUserEmail() {
	user.Email = "theara.dev@gmail.com"
	fmt.Println("User new email:", user.Email)
}

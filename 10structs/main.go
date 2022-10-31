package main

import "fmt"

func main() {

	fmt.Println("Let's go with Structs")

	thereaD := User{"Theara", "theara@gmail.com", true, 18}

	fmt.Println(thereaD)
	fmt.Printf("User Info: %+v\n", thereaD)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

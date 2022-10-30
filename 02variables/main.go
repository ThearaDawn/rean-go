package main

import "fmt"

const LoginToken string = "anytoken"

func main() {
	var fullname string = "Theara"
	fmt.Println(fullname)
	fmt.Printf("Variable is of type: %T \n", fullname)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint16 = 256
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var floatVal float32 = 256.3
	fmt.Println(floatVal)
	fmt.Printf("Variable is of type: %T \n", floatVal)

	// default values and some aliases
	var anotherVar string //int
	fmt.Println(anotherVar)
	fmt.Printf("Variable is of type: %T \n", anotherVar)

	// implicit type

	var website = "learn.io"
	fmt.Println(website)

	// no var style

	numberOfUser := 30000
	fmt.Println(numberOfUser)

	//  const LoginToken
	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)
}

package main

import "fmt"

func main() {

	fmt.Println("Let's go with If Else")

	loginAttempted := 0 // default login attempted 0

	suc := "OK"

	var count int

	if suc == "OK" {
		count = 0
	} else {
		count = loginAttempted + 1
	}

	fmt.Println(count)

	if 6%2 == loginAttempted {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	if num := 9; num < 10 {
		fmt.Println("Number is less than 10")
	} else {
		fmt.Println("Number is larger than 10")
	}

}

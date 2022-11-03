package main

import "fmt"

func main() {
	fmt.Println("Let's go with Function in Golang")
	hello()
	sum(2, 2)
	fmt.Println("Sum two value with return: ", sumReturnValue(3, 3))
}

func hello() {
	fmt.Println("Hello Function")
}

func sum(num1 int, num2 int) {
	fmt.Println("Sum two value: ", num1+num2)
}

func sumReturnValue(num1 int, num2 int) int {
	return num1 + num2
}

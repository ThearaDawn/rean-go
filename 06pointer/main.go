package main

import "fmt"

func main() {

	num := 123
	fmt.Println("Before apply pointer: ", num)

	var pointerRef = &num
	fmt.Println("Pointer value is: ", pointerRef)
	fmt.Println("Actual pointer value is: ", *pointerRef)

	*pointerRef = 2 + *pointerRef
	fmt.Println("New value is: ", num)
	fmt.Println("New pointer value is: ", *pointerRef)
}

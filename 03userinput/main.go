package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our services: ")

	//comma ok OR err

	//input, _ := reader.ReadString('\n')

	input, _ := reader.ReadString('\n')
	fmt.Println("Thank for rating,", input)
	fmt.Printf("Type of this rating is %T", input)

}

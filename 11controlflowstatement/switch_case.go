package main

import "fmt"

func main() {

	fmt.Println("Let's go with Switch Case Statement")

	mDay := 1

	switch mDay {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid num of day")
	}
	// without condition
	switch {
	case 1 == mDay:
		fmt.Println("Monday")
	case 2 == mDay:
		fmt.Println("Tuesday")
	}
	// case list
	switch mDay {
	case 1, 2, 3:
		fmt.Println("Holiday")
	case 4, 5:
		fmt.Println("Working")
	}

}

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Let's go with Map")

	mWeek := make(map[int]string)

	mWeek[0] = "Sunday"
	mWeek[1] = "Monday"
	mWeek[2] = "Tuesday"
	mWeek[3] = "Wednesday"
	mWeek[4] = "Thursday"
	mWeek[5] = "Friday"
	mWeek[6] = "Saturday"

	//fmt.Println("7 Days Of The Week: ", mWeek)
	fmt.Println("Today is", mWeek[1])

	// delete value by key
	delete(mWeek, 0)
	fmt.Println("7 Days Of The Week: ", mWeek)

	// loop map
	/*for key, value := range mWeek {
		fmt.Println("Days of the Week", key, value)
	}*/

	for _, index := range mWeek {
		fmt.Println("Days of the Week:", index)
	}

}

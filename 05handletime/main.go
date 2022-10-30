package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time studing Golang")

	presentTime := time.Now()
	fmt.Println(presentTime)

	//format
	fmt.Println(presentTime.Format("01-02-2006"))

	fmt.Println(presentTime.Format("01-02-2006 Monday"))

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2022, time.October, 30, 8, 52, 0, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}

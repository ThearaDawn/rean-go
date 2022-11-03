package main

import "fmt"

func main() {

	mNums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	for n := 0; n < len(mNums); n++ {
		fmt.Println("Numbers:", mNums[n])
	}

	// range

	for i := range mNums {
		fmt.Println("Numbers =>", mNums[i])
	}

	mValue := 1

	for mValue < 5 {
		if mValue == 3 {
			mValue++ // if not use like this the app will stuck
			//break
			continue
		}
		fmt.Println("M Value is:", mValue)
		mValue++
	}
}

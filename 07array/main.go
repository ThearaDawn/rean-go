package main

import "fmt"

func main() {
	fmt.Println("Rean Arrays")

	var myPlayList [7]string
	var mList = [2]string{"go", "java"}

	myPlayList[0] = "Golang Hello World"
	myPlayList[1] = "Tech"
	myPlayList[2] = "Array in Go"
	myPlayList[3] = "Music"
	myPlayList[4] = "Movie"
	myPlayList[5] = "GoGo"
	myPlayList[6] = "DuckDuck"
	//myPlayList[7] = "Index out of bound"

	fmt.Println("My play list:", myPlayList)
	fmt.Println("Array len", len(myPlayList))

	fmt.Println(mList)

}

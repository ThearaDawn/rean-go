package main

import (
	"fmt"
	"sort"
)

func main() {

	var mList = []string{"Java", "Go", "Python"}

	fmt.Println("My List: ", mList)

	//Append
	mList = append(mList, "PHP", "JS")
	fmt.Println(mList)

	// Remove first element
	// mList = append(mList[1:])

	// Remove first element and all elements after 3
	mList = append(mList[1:3])
	fmt.Println(mList)

	scores := make([]int, 2)
	scores[0] = 123
	scores[1] = 1234

	// Dynamic append
	scores = append(scores, 12345, 123456)
	fmt.Println("Scores: ", scores)

	// Sort
	sort.Ints(scores)
	fmt.Println("Max score:", scores)
}

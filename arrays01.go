package main

import "fmt"

func main() {
	// create an array that holds 5 ints
	var scores [5]int
	fmt.Println("emp:", scores)

	// set a value
	scores[4] = 100
	fmt.Println("set:", scores)
	fmt.Println("get:", scores[4])

	// the builtin len() will return the length
	fmt.Println("len:", len(scores))

	// declare and init in a single line
	highScores := [5]int{1,2,3,4,5}
	fmt.Println("dcl:", highScores)

}

package main

import (
	"fmt"
)

var zzz = "package-level (global) variable\n"

func main() {

	// display our "global" variable
	fmt.Println(zzz)

	// This would be "long hand" decleration
	var mug string = "filled with tea\n"
	fmt.Println(mug)

	// Either type or expression may be omitted (not both)
	// these are all examples of type being ommited
	var a = "Hello"
	var b = 23
	var c = true
	var d = 2.3

	 // var a = "Goodbye"
	a = "Goodbye"
	fmt.Printf("The type of a, b, c, d is: %T, %T, %T and %T respectively\n", a, b, c, d)
     // Prints: The type of a, b, c, d is: string, int, bool and float64 respectively
}

package main

import (
	"fmt"
	"strconv" // built in package for string conversions
)

func main() {

	// create a string via inference
	s := "33"

	// string to int
	i, err := strconv.Atoi(s)
	if err != nil {
		// handle error
		panic(err)
	}
	fmt.Println("String:", s)
	fmt.Println("Integer:", i)

	// transform an int into a string
	st := strconv.Itoa(42)
	fmt.Printf("st is now of type, %T", st)
	fmt.Println("String", st)
}

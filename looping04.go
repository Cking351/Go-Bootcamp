package main

import "fmt"

func main() {

	// create the string alias
	a := []string{"Alta3", "Research", "Loves", "Golang"}

	// The "i" pos is always the index
	// The "s" pos is always the corresponding value
	// The second iteration var is optional
	for i, s := range a {
		fmt.Println("Postition", i, "contains the string:", s)
	}


	for i := range a {
		fmt.Println("Position", i)
	}

	for pos := range a {
		fmt.Println("Position", pos)
	}

}

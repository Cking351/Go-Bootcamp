
package main

import "fmt"

func noPointer() string {
	return "string"
}

func pointerTest() *string {
	return nil	// You can return nil from a function that returns a pointer to a string
}

func pointerTestTwo() *string {
	s := "string"	// &"string" doesn't work
	return &s
}

func main() {
	fmt.Println(noPointer()) // prints string
	fmt.Println(pointerTest()) // returns nil
	fmt.Println(pointerTestTwo()) // prints memory address

	s := pointerTestTwo() // now s holds an address in memory
	fmt.Println(s)
	sp := *s 
	fmt.Println(sp)
}

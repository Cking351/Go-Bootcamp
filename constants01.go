package main

import "fmt"

// constant created in the package level
// constants can appear anywhere a var can appear
const Pi = 3.14

func main() {

	// constant created at a function level
	const MyURI = "http://example.com:5000/v2/"
	fmt.Println("The protocol, auth, and path /v2/ cannot be modified:", MyURI)

	// invoke a variable created a the package level
	fmt.Println("Happy", Pi, "Day")

	// the variable "xfiles" has been set to a bool type
	const Xfiles = true
	fmt.Println("The truth is out there", Xfiles)
}


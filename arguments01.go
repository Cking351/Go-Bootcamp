
package main

import (
	"fmt"
	"os"
)


func main() {

	// os.Args provides access to raw CLI arguments
	argsWithProg := os.Args // includes program name

	argsWithoutProg := os.Args[1:] // does not include program name

	arg := os.Args[3] // selects the third argument

	last_arg := os.Args[len(os.Args)-1] // select the last argument

	fmt.Println("All arguements, including the program name:")
	fmt.Println(argsWithProg)

	fmt.Println("All arguments, without the program name:")
	fmt.Println(argsWithoutProg)

	// select just the 3rd argument
	fmt.Println(arg)

	// display the last argument
	fmt.Println(last_arg)
}

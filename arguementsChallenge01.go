package main

import (
	"os"
	"fmt"
)

func main() {
	argLength := len(os.Args[1:])

	fmt.Printf("Arg length is %d\n", argLength)
}

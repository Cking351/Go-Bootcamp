package main

import (
	"os"
	"fmt"
	"bufio"
)

func main() {
	reader := bufio.NewReader(os.Stdin) // creates a reader obj

	fmt.Print("Enter you name: ")

	name, _ := reader.ReadString('\n') // Read up to the \n character and throw away err var

	fmt.Printf("Hello %s\n", name)
}

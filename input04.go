package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func main() {
	fmt.Print("Enter text: ")
	reader := bufio.NewReader(os.Stdin)

	// readstring will block until the delimter is entered
	input, err := reader.ReadString('\n')

	// error handling
	if err != nil {
		fmt.Println("An error occurred while reading input. Please try again", err)
		return 
	}

	// remove the delimiter from the string
	input = strings.TrimSuffix(input, "\n")
	fmt.Println(input)
}

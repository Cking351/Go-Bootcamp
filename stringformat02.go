package main

import "fmt"

func main() {

	// create some int vars we can use for testing
	n2 := 2
	n3 := 3
	n4 := 4

	res := fmt.Sprintf("There are %d oranges %d apples %d plums", n2, n3, n4)
	fmt.Println(res) // There are 2 oranges 3 apples 4 plums

	// Notice we are not using indexing
	// argument posistion 0 (the string)
	res2 := fmt.Sprintf("There are %[2]d oranges %[3]d apples %[1]d plums", n2, n3, n4)
	fmt.Println(res2)
}

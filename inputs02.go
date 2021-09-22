package main

import "fmt"

func main() {

	var name string
	var age int

	fmt.Print("Enter your name and age: ") // print is not followed by \n
	fmt.Scanf("%s %d", &name, &age)
	fmt.Printf("%s is %d years old\n", name,  age)
}

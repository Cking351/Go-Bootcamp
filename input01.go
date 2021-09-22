package main

import "fmt"

func main() {

	var name string // define a string var for tracking user response

	fmt.Print("Enter your name: ") // this is a comment
	fmt.Scanf("%s", &name)  //need pointer (other we'll make a copy and name change the copy)
	fmt.Println("Hello", name)
}

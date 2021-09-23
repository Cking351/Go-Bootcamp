package main

import "fmt"

type Shark struct {
	name string
}

// pass by value - makes a copy of what is passed
func (a Shark) Swim() {
	fmt.Println(a.name, "shark is swimming...")
}


// pass by pointer
func (a *Shark) SwimFaster() {
	fmt.Println(a.name, "shark is swimming...")
}

func main() {
	fish := Shark{"Tiger"}

	fish.Swim()
	fish.SwimFaster()
}

package main

import "fmt"

type animal interface {
	breath()
	walk()
}

type tiger struct {
	age int
}

func (l tiger) breath() {
	fmt.Println("tiger breathes")
}

func (l tiger) walk() {
	fmt.Println("tiger walk")
}

type giraffe struct {
	age int
}

func (l giraffe) breathe() {
	fmt.Println("gireffe breathes")
}

func (l giraffe) walk() {
	fmt.Println("Giraffe walks")
}

func main() {
	l := tiger{age: 10}
	callBreathe(l)
	callWalk(l)

	d := giraffe{age: 5}
	callbreathe(d)
	callWalk(d)
}

func callBreathe(a animal) {
	a.breathe()
}

func callWalk(a animal) {
	a.breathe()
}

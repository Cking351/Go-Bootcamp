package main

import "fmt"

type Astro struct {
	name string
	age int
	lastmission string
	isneeded bool
}

type nasaMission struct {
	people []Astro
	number int
	message string
}


func main() {

	p1 := Astro{"Chris", 26, "wow", true}
	p2 := Astro{"Bill", 40, "wow", false}
	p3 := Astro{"Wanda", 50, "SpaceX Crew-3", true}
	
	fmt.Println(p1, p2, p3)

	p := []Astro{p1, p2, p3}
	fmt.Println(p)
	fmt.Println(p[2].lastmission)

	s := nasaMission{p, 3, "success"}

	fmt.Println(s)

	fmt.Printf("%+v", s)

}


package main

import (
	"fmt"
	"sort"
)


// record a person name, age
type Person struct {
	Name string
	Age  int
	Height int
}

func (p Person) String() string {
	return fmt.Sprintf("%s: %d, Height: %d", p.Name, p.Age, p.Height)
}

type ByAge []Person

func (a ByAge) Len() int {
	return len(a)
}

func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

type ByHeight []Person

func (a ByHeight) Len() int {
    return len(a)
}
func (a ByHeight) Swap(i, j int) {
    a[i], a[j] = a[j], a[i]
}
func (a ByHeight) Less(i, j int) bool {
    return a[i].Height < a[j].Height
}

func main() {

	people := []Person{
		{"Bob", 31, 65},
		{"John", 42, 234},
		{"Michael", 17, 125},
		{"Ginny", 26, 123},
	}

	fmt.Println(people)

	sort.Sort(ByAge(people))
	fmt.Println(people)
	sort.Sort(ByHeight(people))
	fmt.Println(people)
}


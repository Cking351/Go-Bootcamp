
package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13

	// This will show key value pairs
	fmt.Println("map:", m)

	// return the value associated with "key"
	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	// Determine number of key/value associations
	fmt.Println("len:", len(m))

	// get rid of the "key" k2
	delete(m, "k2")

	// no more key k2
	fmt.Println("map:", m)


	// save only the optional second return value
	// this second value can be used to determine if a key exists
	// or if its missing from the map
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// declare and init a map all in one line
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}

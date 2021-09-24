package main

import (
	"fmt"
	"sort"
)


func main() {

	// create a slice of strings
	strs := []string{"r", "z", "f", "e", "s", "e", "r"}

	// sorting is performed in place. It returns nothing
	sort.Strings(strs)

	fmt.Println("Strings:", strs)

	// create a slice of ints
	ints := []int{2,4,6,8,1}

	sort.Ints(ints)

	fmt.Println("Ints:   ", ints)

	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)
}

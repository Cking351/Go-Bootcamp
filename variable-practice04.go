package main

import (
	"fmt"
)

// t is a package-level variable that can be accessed
// anywhere in the main package. you can say its global to the package
var t = true

func main() {
	// f is block level variable from its point of declaration to the end of the function
	f := false

	{
    // i is a block-level variable that's only valid from this point
    // until the end of the block
        i := 20
        fmt.Println(i) // 20
        } // this block scope is now over so i is no longer valid

    fmt.Println(t, f) // true false
}

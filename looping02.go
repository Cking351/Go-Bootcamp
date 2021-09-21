package main

import (
	"fmt"
	"math/rand"
)

func main() {

	drive := 0 // drive is defined at function level
	fmt.Print("Get a mulligan on any drive under 60 yards.\n")

	// typically after for is the init var, which is omitted
	// no post statement appears the last semi-colon
	for ; drive <= 60; {
		drive = rand.Intn(251)
		fmt.Print("SWING!\n")
	}
	fmt.Println("Your longest drive was", drive, "yards")
}


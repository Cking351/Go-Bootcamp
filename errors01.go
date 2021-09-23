package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("We dont have the power")
	fmt.Println("Scotty says:", err)
}

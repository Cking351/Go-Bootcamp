package main

import (
	"errors"
	"fmt"
)

func rollchar(firstName string, lastName string) (string, error) {
	if lastName == "Turnip" || lastName == "Radish" || lastName == "Potato" {
		return firstName + " the " + lastName, errors.New("Vegetables are not a suitable name")
	}
	return firstName + " the " + lastName, nil
}

func main() {
	fmt.Println("Welcome to the character generator")

	playerChar, err := rollchar("Gandalf", "Turnip")

	if err != nil {
		fmt.Println("Error while spanwing your requested character", err)
	} else {
		fmt.Println(playerChar, "has been generated")
	}
}

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Game: Guess a number between 0 and 10")
	fmt.Println("You have three tries")

	// generating random numbers
	source := rand.NewSource(time.Now().UnixNano())
	// default is not random, it must be seeded
	// not safe for any real secrets

	randomizer := rand.New(source)
	secretNumber := randomizer.Intn(10)

	var guess int
	var rounds int

	fmt.Println("How many rounds?")
	fmt.Scan(&rounds)

	if rounds > 5 {
		fmt.Println("You must now guess a number between 0 and 100")
		secretNumber = randomizer.Intn(100)
	}

	for try := 1; try <= rounds; try++ {
		fmt.Println("Round:", try)

		fmt.Println("Please enter your number")
		fmt.Scan(&guess)

	if guess < secretNumber {
		fmt.Printf("Sorry, number is too small\n")
	} else if guess > secretNumber {
		fmt.Printf("Sorry, number is too high\n")
	} else {
		fmt.Printf("You Win!!\n")
		break
	}

	if try == rounds {
		fmt.Printf("Game over!!\n")
		fmt.Printf("The correct number is %d\n", secretNumber)
		break
	}

	}
	fmt.Println("Thanks for playing!\n")
}

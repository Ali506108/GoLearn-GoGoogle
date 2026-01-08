package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	target := random.Intn(100) + 1

	fmt.Println("Welcome to the guessing game!")

	var guess int

	for {
		fmt.Print("Enter your guess: ")
		fmt.Scanln(&guess)

		if guess == target {
			fmt.Println("Congrats you guesed correct")
			break
		} else if guess < target {
			fmt.Println("your gueses to low")
		} else {
			fmt.Println("Your gueses to high")
		}
	}
}

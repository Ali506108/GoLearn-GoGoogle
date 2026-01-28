package main

import (
	"fmt"
)

func main() {

	fmt.Println("Hello how are you ?")

	fmt.Print("Tell me youre programing language: ")
	var score int
	fmt.Scan(&score)

	if score > 60 {
		fmt.Println("you're engish level is C1 congrats your level is, ", score)
	} else if score <= 60 {
		fmt.Println("Youre Engish levek is B2 , -", score)
	} else if score <= 45 {
		fmt.Println("Youre Engish level is B1 , ", score)
	} else {
		fmt.Println("You're wrong ", score)
	}

	num := 15

	if num%2 == 1 {
		if num%3 == 0 {
			fmt.Println("yeah youre number is ", num)
		} else {
			fmt.Println("normal")
		}
	} else {
		fmt.Println("Wrong")
	}

	if 10%2 == 0 || 6%2 == 0 {
		fmt.Println("Either")
	}

}

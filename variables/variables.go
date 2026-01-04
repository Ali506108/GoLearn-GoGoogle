package main

import (
	"fmt"
)

func main() {
	var name string = "John"
	var age int = 20

	fmt.Printf("tell me you're name %d and age %d ", name, age)

	var uMaxInt uint64 = 4504054050654
	fmt.Println(uMaxInt)

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println()

	numbers := []int{1, 2, 3, 4, 5, 6, 7}

	for index, value := range numbers {
		fmt.Printf("Index %d and value is %d\n", index, value)
	}

	fmt.Println()

	for i := 0; i < 20; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Odd number : ", i)
	}

	fmt.Println()

	rows := 5

	for i := 1; i <= rows; i++ {
		for j := 1; j <= rows-i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

}

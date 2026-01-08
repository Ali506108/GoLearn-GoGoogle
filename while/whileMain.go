package main

import (
	"fmt"
	"log"
)

func main() {

	num := 0

	for num <= 10 {
		if num%2 == 0 {
			log.Default().Println()
			num++
			continue
		}
		fmt.Println(num)
		num++
	}

	Ints := []int{1, 2, 3, 4}
	Strings := []string{"One", "Two", "Three"}
	Print(Ints)
	Print(Strings)
}

func Print[T any](s []T) {
	for _, v := range s {
		fmt.Println(v, " ")
	}
	fmt.Println()
}

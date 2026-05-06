package main

import "fmt"

func main() {

	value := add(5, 5)
	fmt.Println(value)

	if value == 10 {
		func() {
			fmt.Println("Solve the problem")
		}()
	} else {
		fmt.Println("OKAY")
	}

	greet := func() {
		fmt.Println("solve the problem")
	}

	greet()

	fmt.Println(work_with_function(45, 53))

}

func add(value_one int, value_two int) int {
	return value_one + value_two
}

func work_with_function(value_one int, value_two int) int {
	return value_one + value_two*value_two
}

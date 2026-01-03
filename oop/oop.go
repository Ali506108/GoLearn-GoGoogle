package main

import "fmt"

type User struct {
	Name     string
	LastName string
	Age      int
	isOnline bool
}

func main() {
	const MAXRETRIES = 5

	var employeeID = 2
	fmt.Println(employeeID)
}

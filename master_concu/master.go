package main

import (
	"fmt"
	"time"
)

func main() {
	greeting := make(chan string)
	greetString := "Hello world!"

	go func() {
		greeting <- greetString
		greeting <- "World is great"
	}()

	go func() {
		receiver := <-greeting
		fmt.Println("Receiver : ", receiver)
		receiver = <-greeting
		fmt.Println("Second receiver : ", receiver)
	}()

	time.Sleep(1 * time.Second)
}

package main

import (
	"fmt"
	"time"
)

func main() {
	var Ball int
	table := make(chan int)

	go player(table)
	go player(table)
	go player(table)

	table <- Ball
	time.Sleep(1 * time.Second)
	<-table
	fmt.Println("Ball is ", Ball)
}

func player(table chan int) {
	for {
		ball := <-table
		ball++
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Table count is : ", ball)
		table <- ball
	}
}

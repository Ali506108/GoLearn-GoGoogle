package main

import (
	"fmt"
	"time"
)

// Fan-in это функция, читающая из нескольких источников и мультиплексирующая всё в один канал.
func producer(ch chan int, d time.Duration) {
	var i int
	for {
		ch <- i
		i++
		time.Sleep(d)
	}
}

func reader(out chan int) {
	for x := range out {
		fmt.Println(x)
	}
}

func main() {
	prod := make(chan int)
	out := make(chan int)

	go producer(prod, 1000*time.Microsecond)
	go producer(prod, 1000*time.Millisecond)

	go reader(out)

	for i := range prod {
		out <- i
	}
}

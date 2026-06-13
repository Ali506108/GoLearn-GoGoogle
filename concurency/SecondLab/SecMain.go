package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	for i := 1; i <= 5; i++ {
	// 		ch <- i
	// 	}
	// 	close(ch)
	// }()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for val := range ch {
			fmt.Println("Got: ", val)
		}
	}()

	wg.Wait()
}

package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	counter := 0

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			fmt.Printf("%d\n", i)
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println("Final counter: ", counter)
}

package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	fmt.Printf("%#v\n", wg)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			fmt.Printf("%d", x)
		}(i)
	}
	fmt.Printf("%#v\n", wg)
	wg.Wait()
}

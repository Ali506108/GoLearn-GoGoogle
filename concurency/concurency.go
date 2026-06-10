package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func handleSignal(sig os.Signal) {
	fmt.Println("HandelSignal() Caught: ", sig)
}

func main() {
	fmt.Printf("Process ID : $d\n", os.Getpid())
	fmt.Println("Process Id : $d\n", os.Getpid())

	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs)

	start := time.Now()

	go func() {
		for {
			sig := <-sigs

			switch sig {
			case syscall.SIGINT:
				duration := time.Since(start)
				fmt.Println("Duration : ", duration)
			case syscall.SIGINFO:
				handleSignal(sig)

			}
		}
	}()
}

package main

import (
	"fmt"
	"time"
)

func main() {
	// goroutines
	for i := 0; i < 3; i++ {
		go func() {
			fmt.Println("goroutines is how you achieve concurrency in Go.")
		}()
	}
	time.Sleep(1 * time.Second)

	// Go channels
	fmt.Println("Go channels allow you to send data across goroutines")
	fmt.Println("Do not communicate by sharing memory; instead, share memory by communicating.")
	queue := make(chan int)
	go Work(queue)
	queue <- 1
	queue <- 2
	queue <- 3
	time.Sleep(1 * time.Second)
}

func Work(queue chan int) {
	for work := range queue {
		fmt.Printf("Working with %d...\n", work)
	}
}

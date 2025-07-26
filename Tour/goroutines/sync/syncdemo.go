package main

import (
	"fmt"
	"sync"
	"time" // For demonstration, to exaggerate scheduling differences
)

func sum(s []int, c chan int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the counter when the goroutine finishes

	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
	fmt.Println("Call completed after received")
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	var wg sync.WaitGroup // Create a WaitGroup

	// Add 2 to the WaitGroup counter, one for each goroutine
	wg.Add(2)

	go sum(s[:len(s)/2], c, &wg)
	go sum(s[len(s)/2:], c, &wg)

	// Receive from the channel
	x := <-c
	y := <-c

	fmt.Println("Main goroutine received values, waiting for others to complete...")
	time.Sleep(10 * time.Millisecond) // Small sleep to allow scheduler to run the print, if needed
	// (WaitGroup is the real solution, but this highlights scheduling)

	// Wait for both goroutines to call wg.Done()
	wg.Wait()

	fmt.Println(x, y, x+y)
}

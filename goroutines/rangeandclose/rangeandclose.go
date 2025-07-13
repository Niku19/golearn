package main

import (
	"fmt"
)

// A sender can close a channel to indicate that no more values will be sent. Receivers can test whether a channel has been closed by assigning a second parameter to the receive expression: after

// v, ok := <-ch
// ok is false if there are no more values to receive and the channel is closed.

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for range n {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

// The loop for i := range c receives values from the channel repeatedly until it is closed.
// Sending on a closed channel will cause a panic.
// Channels aren't like files; you don't usually need to close them.
// Closing is only necessary when the receiver must be told there are no more values coming, such as to terminate a range loop.
func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

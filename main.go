package main

import "fmt"

func main() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3

	// Receive values from the channel
	val1 := <-ch
	val2 := <-ch
	val3 := <-ch

	fmt.Println(val1, val2, val3) // Output: 1 2 3

	// The channel is still open and not cleared
	fmt.Println("Channel length:", len(ch)) // Output: Channel length: 0

	// Try to receive from the channel again
	val4, ok := <-ch
	if !ok {
		fmt.Println("Channel is closed")
	} else {
		fmt.Println("Received value:", val4)
	}
}

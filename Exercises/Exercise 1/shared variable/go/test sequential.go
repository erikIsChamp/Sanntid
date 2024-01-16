package main

import (
	"fmt"
	"time"
)

func incrementAndDecrement() int {
	var sharedValue int

	// Increment 1,000,000 times
	for j := 0; j < 1000000; j++ {
		sharedValue++
	}

	// Decrement 1,000,000 times
	for j := 0; j < 1000000; j++ {
		sharedValue--
	}

	return sharedValue
}

func main() {
	// Start the timer
	startTime := time.Now()

	// Perform increment and decrement operations sequentially
	result := incrementAndDecrement()

	// Stop the timer
	endTime := time.Now()

	// Calculate and print the runtime
	runtime := endTime.Sub(startTime)

	fmt.Println("The magic number is:", result)
	fmt.Printf("Runtime: %v\n", runtime)
}

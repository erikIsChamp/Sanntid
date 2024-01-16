package main

import (
	"fmt"
	"sync"
	"time"
)

// incrementer increments a shared integer and signals completion.
func incrementer(incrementCh chan int, doneCh chan bool, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the WaitGroup counter when done
	for j := 0; j < 1000000; j++ {
		incrementCh <- 1 // Signal to increment
	}
}

// decrementer decrements a shared integer and signals completion.
func decrementer(decrementCh chan int, doneCh chan bool, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the WaitGroup counter when done
	for j := 0; j < 1000000; j++ {
		decrementCh <- 1 // Signal to decrement
	}
}

func main() {
	var sharedValue int
	incrementCh := make(chan int)
	decrementCh := make(chan int)
	doneCh := make(chan bool)

	var wg sync.WaitGroup

	// Add two goroutines to the WaitGroup
	wg.Add(2)

	// Start the timer
	startTime := time.Now()

	// Start the incrementer and decrementer goroutines
	go incrementer(incrementCh, doneCh, &wg)
	go decrementer(decrementCh, doneCh, &wg)

	// Create a goroutine to manage the shared integer based on signals
	go func() {
		for {
			select {
			case <-incrementCh:
				continue
			case <-decrementCh:
				continue
			}
		}
	}()

	// Wait for both incrementer and decrementer goroutines to finish
	wg.Wait()

	// Close the doneCh to signal completion
	close(doneCh)

	// Stop the timer
	endTime := time.Now()

	// Calculate and print the runtime
	runtime := endTime.Sub(startTime)
	fmt.Println("The magic number is:", sharedValue)
	fmt.Printf("Runtime: %v\n", runtime)
}

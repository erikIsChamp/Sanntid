package main

import (
	"fmt"
	"sync"
)

func incrementer(incrementCh chan int, doneCh chan bool, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the WaitGroup counter when done
	for j := 0; j < 1000000; j++ {
		incrementCh <- 1 // Signal to increment
	}
}

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

	go incrementer(incrementCh, doneCh, &wg)
	go decrementer(decrementCh, doneCh, &wg)

	// Create a goroutine to manage the shared integer based on signals
	go func() {
		for {
			select {
			case <-incrementCh:
				sharedValue++
			case <-decrementCh:
				sharedValue--
			}
		}
	}()

	// Wait for both incrementer and decrementer goroutines to finish
	wg.Wait()

	// Close the doneCh to signal completion
	close(doneCh)

	fmt.Println("The magic number is:", sharedValue)
}

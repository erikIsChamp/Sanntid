// Use `go run foo.go` to run your program

package main

import (
	"fmt"
	. "fmt"
	"runtime"
)

var i = 0

func incrementing(c chan int) {
	// TODO: increment i 1000000 times
	for j := 0; j <= 1000000; j++ {
		i++
	}
	c <- i // Send a signal to the channel

	close(c) // Close the channel when done
}

func decrementing(c chan int) {
	//TODO: decrement i 1000000 times
	for j := 0; j <= 1000000; j++ {
		i--
	}
	c <- i // Send a signal to the channel

	close(c) // Close the channel when done
}

func main() {
	// What does GOMAXPROCS do? What happens if you set it to 1?
	// It enables the use of 2 threads to be used by Go. If you set it to 1, only 1 thread can be used
	runtime.GOMAXPROCS(2)

	channel := make(chan int)

	// TODO: Spawn both functions as goroutines

	// Start function1 as a goroutine
	go incrementing(channel)

	// Start function2 as a goroutine
	go decrementing(channel)

	// Use select to receive signals from the goroutines
	for {
		select {
		case signal, ok := <-channel:
			if !ok {
				// Channel is closed, and all goroutines are done
				fmt.Println("Both functions have finished.")
				Println("The magic number is:", i)
				return
			}
			// Handle the signal received from the goroutines
			if signal == 1 {
				fmt.Println("Received signal from function1")
			} else if signal == 2 {
				fmt.Println("Received signal from function2")
			}
		}
	}

}

// Use `go run foo.go` to run your program

package main

import (
	"fmt"
	. "fmt"
	"runtime"
	"time"
)

var i = 0

func incrementing() {
	// TODO: increment i 1000000 times
	for j := 0; j <= 1000000; j++ {
		i++
	}
}

func decrementing() {
	//TODO: decrement i 1000000 times
	for j := 0; j <= 1000000; j++ {
		i--
	}
}

func main() {
	// What does GOMAXPROCS do? What happens if you set it to 1?
	// It enables the use of 2 threads to be used by Go. If you set it to 1, only 1 thread can be used
	runtime.GOMAXPROCS(2)

	// TODO: Spawn both functions as goroutines

	// Start the first goroutine
	go func() {
		incrementing()
	}()

	// Start the second goroutine
	go func() {
		decrementing()
	}()

	fmt.Println("Both goroutines have completed.")

	// We have no direct way to wait for the completion of a goroutine (without additional synchronization of some sort)
	// We will do it properly with channels soon. For now: Sleep.
	time.Sleep(500 * time.Millisecond)
	Println("The magic number is:", i)
}

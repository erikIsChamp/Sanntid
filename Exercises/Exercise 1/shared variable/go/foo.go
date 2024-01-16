
// Use go run foo.go to run your program

package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	i            = 0
	incrementCh  = make(chan struct{})
	decrementCh  = make(chan struct{})
	getValueCh   = make(chan int)
	workerFinish = make(chan struct{})
)

// NumberServer represents the server that manages the integer value.
func NumberServer() {
	for {
		select {
		case <-incrementCh:
			i++
		case <-decrementCh:
			i--
		case getValueCh <- i:
			// Do nothing for "get" as the result will be read later
		case <-workerFinish:
			return
		}
	}
}

func incrementing() {
	for j := 0; j < 1000000; j++ {
		incrementCh <- struct{}{}
	}
	workerFinish <- struct{}{}
}

func decrementing() {
	for j := 0; j < 1000006; j++ {
		decrementCh <- struct{}{}
	}
	workerFinish <- struct{}{}
}

func main() {
	runtime.GOMAXPROCS(2)

	go NumberServer()
	go incrementing()
	go decrementing()

	// Wait for both workers to finish
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		<-workerFinish
	}()

	go func() {
		defer wg.Done()
		<-workerFinish
	}()

	// Wait for all goroutines to finish
	wg.Wait()

	// Get the final value from the NumberServer
	finalValue := <-getValueCh

	fmt.Println("The magic number is:", finalValue)
}
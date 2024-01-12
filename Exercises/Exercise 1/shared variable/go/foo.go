// Use `go run foo.go` to run your program

package main

import (
    . "fmt"
    "runtime"
    "time"
    
)

var i int = 0

func incrementing(inc chan int, c chan bool) {
    //TODO: increment i 1000000 times
    for j := 0; j < 1000000; j++ {
        
        i++
        inc <- i
    
    }
    c <- true
}

func decrementing(dec chan int, c chan bool) {
    //TODO: decrement i 1000000 times
    for j := 0; j < 1000001; j++ {
        i--
        dec <- i
    }
    c <- true
}

func main() {
    // What does GOMAXPROCS do? What happens if you set it to 1?
    runtime.GOMAXPROCS(2)
    inc := make(chan int,1)
    dec := make(chan int,1)
    done := make(chan bool,2)
    

	
    // TODO: Spawn both functions as goroutines
    go incrementing(inc,done)
    go decrementing(dec,done)
    
    
        select {   
        case <- inc:
        case <- dec:
        case <- done:
        case <- done:
        }
    
	
    // We have no direct way to wait for the completion of a goroutine (without additional synchronization of some sort)
    // We will do it properly with channels soon. For now: Sleep.
    time.Sleep(500*time.Millisecond)
    Println("The magic number is:", i)
}



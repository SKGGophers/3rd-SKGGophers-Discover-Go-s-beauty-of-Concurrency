package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func init() {
	// Downgrading the number of available threads to avoid parallelism
	// Single-thread go program in other words
	runtime.GOMAXPROCS(1)
}

// Wait for task pattern ( used in pooling implementation patterns )
func main() {
	before := time.Now()
	items := 1000
	// Buffered channel init
	ch := make(chan string, items)

	for i := 0; i < items; i++ {
		// Don't do this in production - it's for simulation/demo purposes ONLY!!
		go func(item int) {
			time.Sleep(time.Duration(rand.Intn(5)) + time.Second)
			ch <- fmt.Sprintf("ok %d", item)
			fmt.Printf("`ok``signal sent for item %d!\n", item)
		}(i)
	}

	for items > 0 {
		val := <-ch
		items--
		fmt.Printf("signal received `%s\n", val)
	}
	close(ch)

	after := time.Now()
	fmt.Printf("exitting main program. time elapsed : %v\n", after.Sub(before))
}

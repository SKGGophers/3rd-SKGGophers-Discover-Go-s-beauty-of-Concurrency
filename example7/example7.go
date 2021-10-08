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
	runtime.GOMAXPROCS(2)
}

// Wait for task pattern ( used in pooling implementation patterns )
func main() {
	before := time.Now()
	// Unbuffered channel init
	ch := make(chan string)
	go func() {
		for {
			println("receiving signal...")
			val, ok := <-ch
			if !ok {
				break
			}
			fmt.Printf("`%s` signal received!\n\n\n", val)
		}

	}()
	for i := 0; i < 10; i++ {
		// Don't do this in production - it's for simulation/demo purposes ONLY!!
		time.Sleep(time.Duration(rand.Intn(600)) + time.Millisecond)
		println("sending signal...")
		ch <- "ok"
		println("`ok``signal sent!")
	}
	close(ch)

	after := time.Now()
	fmt.Printf("exitting main program. time elapsed : %v\n", after.Sub(before))
}

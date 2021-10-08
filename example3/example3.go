package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var counter int

func init() {
	// Downgrading the number of available threads to avoid parallelism
	// Single-thread go program in other words
	runtime.GOMAXPROCS(2)
}

func main() {
	before := time.Now()
	var wg sync.WaitGroup
	wg.Add(2)

	// Creating and executing an un-named / literal function
	go func() {
		add()
		wg.Done()
	}()

	go func() {
		substract()
		wg.Done()
	}()

	// Blocks the main go-routine and waits all goroutines registered under the same wait group to finish
	// before terminating the main go routine
	wg.Wait()
	fmt.Printf("counter value is %d\n", counter)
	after := time.Now()
	fmt.Printf("exitting main program. time elapsed : %v\n", after.Sub(before))
}

func add() {
	for count := 0; count < 1000000; count++ {
		value := counter
		//fmt.Printf("current counter %d\n", counter)
		value++
		counter = value
	}
}

func substract() {
	for count := 0; count < 1000000; count++ {
		value := counter
		//fmt.Printf("current counter %d\n", counter)
		value--
		counter = value
	}
}

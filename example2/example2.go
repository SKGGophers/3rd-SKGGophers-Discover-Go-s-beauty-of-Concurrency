package main

import (
	"fmt"
	"runtime"
	"sync"
)

func init() {
	// Downgrading the number of available threads to avoid parallelism
	// Single-thread go program in other words
	runtime.GOMAXPROCS(1)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	// Creating and executing an un-named / literal function
	go func() {
		englishAlphabet()
		wg.Done()
	}()
	go func() {
		greekAlphabet()
		wg.Done()
	}()


	// Blocks the main go-routine and waits all goroutines registered under the same wait group to finish
	// before terminating the main go routine
	wg.Wait()
	println("exitting main program")
}

func englishAlphabet() {
	println("printing the english alphabet")
	abc := "ABCDEFGHIJKLMNOPQRSTUVWYXZ"
	for _, l := range abc {
		fmt.Printf("%s ", string(l))
	}
	println()
}


func greekAlphabet() {
	println("printing the greek alphabet")
	abc := "ΑΒΓΔΕΖΗΘΙΚΛΜΝΞΟΠΡΣΤΥΦΧΨΩ"
	for _, l := range abc {
		fmt.Printf("%s ", string(l))
	}
	println()
}
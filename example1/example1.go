package main

import (
	"runtime"
	"time"
)

func init() {
	// Downgrading the number of available threads to avoid parallelism
	// Single-thread go program in other words
	runtime.GOMAXPROCS(1)
}

func main() {
	go func() {
		println("running inside the go routing")
	}()

	time.Sleep(time.Second)
	println("running in main go thread")
}

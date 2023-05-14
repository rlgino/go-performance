package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/debug"
	"runtime/trace"

	"go-performance/fibonacci"
)

func main() {
	//The execution tracer in Go provides detailed information about what’s happening in a program, including stack traces, goroutine blocking, and more.
	f, err := os.Create("trace.out")
	if err != nil {
		panic("File not created: " + err.Error())
	}
	if err := trace.Start(f); err != nil {
		panic("Error starting tracer: " + err.Error())
	}
	// Set the max number of processors to use
	runtime.GOMAXPROCS(2)
	// Set the minimum heap size at 1GB
	runtime.MemProfileRate = 1 << 30
	//Set the percentage of Garbage collector
	debug.SetGCPercent(50)
	fmt.Printf("The fibonacci for 20 value is %d\n", fibonacci.Calculate(20))
}

package fibonacci_test

import (
	"go-performance/fibonacci"
	"testing"
)

// Writing effective benchmarks in Go is crucial for understanding the performance of your code.
// go test -bench=. -benchmem
// to generate a file to benchmem
func BenchmarkCalculate(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		fibonacci.Calculate(20)
	}
}

// Go has built-in profiling tools that can help you gain insight into what your code is doing.
// * For generating file.out with data: go test -cpuprofile=prof.out
// * To watch the results: go tool pprof prof.out
// This will open the pprof interactive shell, where we can type various commands
// to analyze the profiling data.
// For example, we can use the “top” command to display the functions that consumed the most CPU time:
func TestCalculate(t *testing.T) {
	result := fibonacci.Calculate(20)
	if result != 6765 {
		t.Errorf("Expected result %d distinct of result %d", 6765, result)
	}
}

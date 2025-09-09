# Go Advanced Testing: Table-Driven Tests and Benchmarking

Go's `testing` package provides powerful features for writing robust tests and measuring performance. This tutorial focuses on two advanced aspects:
1.  **Table-Driven Tests**: An idiomatic way to test multiple scenarios for a function with a single test function.
2.  **Benchmarking**: Measuring the performance of your code.

## 1. Table-Driven Tests

Table-driven tests are a common pattern in Go for testing a function with various inputs and expected outputs. They make your tests concise, readable, and easy to extend. You define a slice of structs, where each struct represents a test case with input parameters, expected results, and a descriptive name.

### Example: `TestSum`

In `21_Testing_Advanced_test.go`, the `TestSum` function uses a table-driven approach to test the `Sum` function with different combinations of positive, negative, and zero inputs.

```go
package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		expected int
	}{
		{"positive numbers", 1, 2, 3},
		{"negative numbers", -1, -2, -3},
		{"mixed numbers", -1, 2, 1},
		{"zero", 0, 0, 0},
		{"large numbers", 1000, 2000, 3000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Sum(tt.a, tt.b)
			if actual != tt.expected {
				t.Errorf("Sum(%d, %d): expected %d, got %d", tt.a, tt.b, tt.expected, actual)
			}
		})
	}
}
```

## 2. Benchmarking

Go's testing package also supports writing benchmarks to measure the performance of your code. Benchmark functions start with `Benchmark` and take a `*testing.B` parameter. Inside the benchmark function, you typically run the code you want to measure `b.N` times.

### Example: `BenchmarkFactorial`

In `21_Testing_Advanced_test.go`, the `BenchmarkFactorial` function measures the performance of the `Factorial` function.

```go
package main

import (
	"testing"
)

func BenchmarkFactorial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Factorial(10) // Benchmark Factorial of 10
	}
}
```

## Running the Examples

Navigate to the `21 Testing Advanced` directory in your terminal.

### Running Tests

To run the tests, execute:

```bash
go test
```

This will run the `TestSum` function and report the results.

### Running Benchmarks

To run the benchmarks, execute:

```bash
go test -bench=.
```

This will run all benchmark functions in the package and report their performance metrics (e.g., operations per second, time per operation).

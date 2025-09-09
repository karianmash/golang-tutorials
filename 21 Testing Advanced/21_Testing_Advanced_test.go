package main

import (
	"testing"
)

// Table-Driven Test for Sum function
func TestSum(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
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

// Benchmark for Factorial function
func BenchmarkFactorial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Factorial(10) // Benchmark Factorial of 10
	}
}

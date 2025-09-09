package main

// Sum calculates the sum of two integers.
func Sum(a, b int) int {
	return a + b
}

// Factorial calculates the factorial of a non-negative integer.
func Factorial(n int) int {
	if n < 0 {
		return 0 // Or handle error, for simplicity returning 0
	}
	if n == 0 {
		return 1
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

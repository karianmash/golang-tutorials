package main

import "fmt"

// IntMin returns the smaller of two integers.
func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(IntMin(1, 2))
}

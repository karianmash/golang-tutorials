package main

import "fmt"

// Functions are a fundamental building block in Go.
// They allow you to group a sequence of statements to perform a specific task.

// A simple function that takes two integers and returns their sum.
// The type of the return value is specified after the parameter list.
func plus(a int, b int) int {
	return a + b
}

// When you have multiple consecutive parameters of the same type,
// you can omit the type name for the like-typed parameters up to the final one.
func plusPlus(a, b, c int) int {
	return a + b + c
}

// Functions can return multiple values.
// This is a common feature in Go, especially for returning an error value along with the result.
func vals() (int, int) {
	return 3, 7
}

// Variadic functions can be called with any number of trailing arguments.
// For example, `fmt.Println` is a common variadic function.
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	// Calling a function is straightforward.
	res := plus(1, 2)
	fmt.Println("1+2 =", res) // Output: 1+2 = 3

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res) // Output: 1+2+3 = 6

	// Here we use the two different return values from the `vals` function.
	a, b := vals()
	fmt.Println(a) // Output: 3
	fmt.Println(b) // Output: 7

	// If you only want a subset of the returned values, you can use the blank identifier `_`.
	_, c := vals()
	fmt.Println(c) // Output: 7

	// Variadic functions can be called in the usual way with individual arguments.
	sum(1, 2)
	sum(1, 2, 3)

	// If you already have multiple args in a slice, you can apply them to a variadic
	// function using the `...` syntax.
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}

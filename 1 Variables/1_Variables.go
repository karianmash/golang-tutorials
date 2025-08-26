package main

import "fmt"

func main() {
	// Go is a statically typed language. This means that variable types are known at compile time.
	// There are several ways to declare variables in Go.

	// 1. Using the 'var' keyword, specifying the type.
	// This declares a variable named 'age' of type 'int' (integer).
	// If not initialized, it will have a zero value for its type (0 for int).
	var age int
	fmt.Println("My age is:", age) // Output: My age is: 0

	// You can also initialize the variable at the time of declaration.
	var name string = "Alice"
	fmt.Println("My name is:", name) // Output: My name is: Alice

	// 2. Type inference.
	// Go can infer the type of the variable from the value it is initialized with.
	var city = "New York"
	fmt.Println("I live in:", city) // Output: I live in: New York

	// 3. Short variable declaration.
	// This is the most common way to declare and initialize variables in Go.
	// It can only be used inside functions.
	country := "USA"
	fmt.Println("I am from:", country) // Output: I am from: USA

	// Multiple variable declarations
	var a, b, c int = 1, 2, 3
	fmt.Println("a, b, c:", a, b, c) // Output: a, b, c: 1 2 3

	// You can also declare variables of different types in a single line.
	var (
		firstName = "John"
		lastName  = "Doe"
		height    = 1.80
	)
	fmt.Println("My name is", firstName, lastName, "and my height is", height)

	// Basic data types in Go
	var integer int = 10
	var float float64 = 3.14
	var boolean bool = true
	var text string = "Hello, Go!"

	fmt.Println("Integer:", integer)
	fmt.Println("Float:", float)
	fmt.Println("Boolean:", boolean)
	fmt.Println("String:", text)
}

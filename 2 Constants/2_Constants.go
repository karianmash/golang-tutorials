package main

import "fmt"

// Constants are declared like variables, but with the 'const' keyword.
// They can be of any basic data type.
const Pi = 3.14

// Constants can be character, string, boolean, or numeric values.
// They cannot be changed once they are declared.

func main() {
	const World = "World"
	fmt.Println("Hello", World)     // Output: Hello World
	fmt.Println("Happy", Pi, "Day") // Output: Happy 3.14 Day

	// This will cause a compile-time error because you cannot change the value of a constant.
	// Pi = 3.14159

	// Typed constants
	const Truth bool = true
	fmt.Println("Go rules?", Truth) // Output: Go rules? true

	// Untyped constants
	// Untyped constants take the type needed by their context.
	const ( // You can declare multiple constants in a block
		Big   = 1 << 100
		Small = Big >> 99
	)

	// This function needs an int64
	fmt.Println(needInt(Small)) // Output: 2

	// This function needs a float64
	fmt.Println(needFloat(Small)) // Output: 2
	fmt.Println(needFloat(Big))   // Output: 1.2676506002282295e+30
}

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}


package main

import "fmt"

// A struct is a collection of fields.
// They are useful for grouping data together to form records.

type person struct {
	name string
	age  int
}

func main() {
	// This syntax creates a new struct.
	fmt.Println(person{"Bob", 20}) // Output: {Bob 20}

	// You can name the fields when initializing a struct.
	fmt.Println(person{name: "Alice", age: 30}) // Output: {Alice 30}

	// Omitted fields will be zero-valued.
	fmt.Println(person{name: "Fred"}) // Output: {Fred 0}

	// An `&` prefix yields a pointer to the struct.
	fmt.Println(&person{name: "Ann", age: 40}) // Output: &{Ann 40}

	// Access struct fields with a dot.
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name) // Output: Sean

	// You can also use dots with struct pointers - the pointers are automatically dereferenced.
	sp := &s
	fmt.Println(sp.age) // Output: 50

	// Structs are mutable.
	sp.age = 51
	fmt.Println(sp.age) // Output: 51
}

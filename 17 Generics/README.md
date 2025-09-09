# Go Generics Tutorial

This tutorial introduces Go's generics feature, allowing you to write flexible and reusable functions and data structures that work with any type.

## What are Generics?

Generics enable you to define functions and types with type parameters. This means you can write a single function or type definition that operates on a variety of types, rather than writing separate implementations for each type.

## Example: Generic Print Function

The `17_Generics.go` file contains an example of a generic function `printSlice` that can print elements of any slice type.

```go
package main

import "fmt"

// Generic function to print any slice
func printSlice[T any](s []T) {
	for _, v := range s {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

func main() {
	fmt.Println("Go Generics Tutorial")

	// Example with int slice
	intSlice := []int{1, 2, 3, 4, 5}
	fmt.Print("Int Slice: ")
	printSlice(intSlice)

	// Example with string slice
	stringSlice := []string{"hello", "world", "generics"}
	fmt.Print("String Slice: ")
	printSlice(stringSlice)
}
```

## Running the Example

To run this example, navigate to the `17 Generics` directory in your terminal and execute:

```bash
go run 17_Generics.go
```

This will output:

```
Go Generics Tutorial
Int Slice: 1 2 3 4 5 
String Slice: hello world generics 
```

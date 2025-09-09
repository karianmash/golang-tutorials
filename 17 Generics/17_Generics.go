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

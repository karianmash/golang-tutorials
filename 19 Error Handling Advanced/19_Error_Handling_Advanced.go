package main

import (
	"errors"
	"fmt"
)

// 1. Custom Error Type
// Define a custom error type by implementing the error interface.
type MyError struct {
	Code    int
	Message string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

// 2. Function returning a custom error
func doSomething(value int) error {
	if value < 0 {
		return &MyError{Code: 1001, Message: "Value cannot be negative"}
	}
	return nil
}

// 3. Error Wrapping with %w
// This function simulates a deeper operation that might return an error.
// It wraps the underlying error to add context.
var ErrDatabase = errors.New("database operation failed")

func fetchDataFromDB(query string) error {
	// Simulate a database error
	if query == "invalid" {
		return fmt.Errorf("failed to execute query: %w", ErrDatabase)
	}
	return nil
}

// 4. Inspecting Wrapped Errors with errors.Is and errors.As
// errors.Is checks if an error in the chain matches a specific error.
// errors.As extracts a specific error type from an error chain.

func processData(data string) error {
	if err := fetchDataFromDB(data); err != nil {
		// Check if the underlying error is ErrDatabase
		if errors.Is(err, ErrDatabase) {
			return fmt.Errorf("data processing failed due to database issue: %w", err)
		}
		return fmt.Errorf("data processing failed: %w", err)
	}

	// Simulate another error that might be a custom type
	if data == "error_value" {
		return &MyError{Code: 1002, Message: "Specific processing error"}
	}

	return nil
}

func main() {
	fmt.Println("--- Custom Errors ---")
	if err := doSomething(-5); err != nil {
		fmt.Println("Received error:", err)
		var myErr *MyError
		if errors.As(err, &myErr) {
			fmt.Printf("It's a MyError! Code: %d, Message: %s\n", myErr.Code, myErr.Message)
		}
	}

	fmt.Println("\n--- Error Wrapping and Inspection ---")

	// Case 1: Database error wrapped
	if err := processData("invalid"); err != nil {
		fmt.Println("Received error:", err)
		if errors.Is(err, ErrDatabase) {
			fmt.Println("Underlying error is ErrDatabase.")
		}
	}

	// Case 2: Custom error from processData
	if err := processData("error_value"); err != nil {
		fmt.Println("Received error:", err)
		var myErr *MyError
		if errors.As(err, &myErr) {
			fmt.Printf("It's a MyError! Code: %d, Message: %s\n", myErr.Code, myErr.Message)
		}
	}

	// Case 3: No error
	if err := processData("valid"); err != nil {
		fmt.Println("Received error:", err)
	} else {
		fmt.Println("No error for valid data.")
	}
}

# Go Advanced Error Handling

Go's approach to error handling is explicit and idiomatic. While basic error handling involves checking `err != nil`, advanced scenarios require understanding custom error types, error wrapping, and how to inspect error chains.

This tutorial covers:
1.  **Custom Error Types**: Defining your own error types that implement the `error` interface.
2.  **Error Wrapping (`fmt.Errorf` with `%w`)**: Adding context to errors while preserving the original (wrapped) error.
3.  **Inspecting Wrapped Errors (`errors.Is` and `errors.As`)**: How to check for specific errors within a chain and extract custom error types.

## 1. Custom Error Types

You can define your own error types by creating a struct that implements the `Error() string` method. This allows you to attach additional information (like error codes or specific details) to your errors.

```go
type MyError struct {
	Code    int
	Message string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}
```

## 2. Error Wrapping with `%w`

Go 1.13 introduced error wrapping, allowing you to create a chain of errors. The `%w` verb in `fmt.Errorf` wraps an error, making it retrievable later. This is crucial for adding context as an error propagates up the call stack without losing the original cause.

```go
var ErrDatabase = errors.New("database operation failed")

func fetchDataFromDB(query string) error {
	if query == "invalid" {
		return fmt.Errorf("failed to execute query: %w", ErrDatabase)
	}
	return nil
}
```

## 3. Inspecting Wrapped Errors with `errors.Is` and `errors.As`

-   `errors.Is(err, target)`: Checks if `err` or any error in its chain matches `target`.
-   `errors.As(err, &target)`: Finds the first error in `err`'s chain that matches the type of `target` and assigns it to `target`.

These functions are essential for robust error handling, allowing you to react to specific underlying errors regardless of how many layers of wrapping they have.

## Example: `19_Error_Handling_Advanced.go`

The example file demonstrates all these concepts:
-   `MyError` as a custom error type.
-   `doSomething` returns `MyError`.
-   `fetchDataFromDB` returns a wrapped `ErrDatabase`.
-   `processData` uses `errors.Is` to check for `ErrDatabase` and `errors.As` to extract `MyError`.

```go
package main

import (
	"errors"
	"fmt"
)

// ... (code as in 19_Error_Handling_Advanced.go)

func main() {
	// ... (main function as in 19_Error_Handling_Advanced.go)
}
```

## Running the Example

To run this example, navigate to the `19 Error Handling Advanced` directory in your terminal and execute:

```bash
go run 19_Error_Handling_Advanced.go
```

Observe how the program identifies and reacts to different types of errors, including custom errors and wrapped errors, demonstrating the power of `errors.Is` and `errors.As`.

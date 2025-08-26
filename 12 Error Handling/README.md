
# 12. Error Handling in Go

This section explains how to handle errors in Go.

## Returning Errors

In Go, it is idiomatic to handle errors by returning an `error` value as the last return value of a function.

```go
func f1(arg int) (int, error) {
    if arg == 42 {
        return -1, errors.New("can't work with 42")
    }
    return arg + 3, nil
}
```

- If the function completes successfully, it returns a `nil` error.
- If the function encounters an error, it returns an `error` object with a message describing the error.

## Handling Errors

When you call a function that returns an error, you should check if the error is `nil`.

```go
if r, e := f1(i); e != nil {
    fmt.Println("f1 failed:", e)
} else {
    fmt.Println("f1 worked:", r)
}
```

This is a common pattern in Go for handling errors.

## Custom Error Types

You can create your own custom error types by implementing the `Error()` method.

```go
type argError struct {
    arg  int
    prob string
}

func (e *argError) Error() string {
    return fmt.Sprintf("%d - %s", e.arg, e.prob)
}
```

This allows you to create more structured errors that can be inspected programmatically.

To check for a specific type of error, you can use a type assertion.

```go
_, e := f2(42)
if ae, ok := e.(*argError); ok {
    fmt.Println(ae.arg)
    fmt.Println(ae.prob)
}
```

## Running the Code

To run the code in `12_Error_Handling.go`, use the following command:

```bash
go run "12 Error Handling/12_Error_Handling.go"
```

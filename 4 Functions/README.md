
# 4. Functions in Go

This section covers the basics of creating and using functions in Go.

## Defining Functions

A function is a block of code that performs a specific task. You can define a function using the `func` keyword.

```go
func functionName(parameters) returnType {
    // function body
}
```

- `functionName`: The name of the function.
- `parameters`: The input parameters for the function.
- `returnType`: The type of the value that the function returns.

Here is a simple example:

```go
func plus(a int, b int) int {
    return a + b
}
```

If you have multiple parameters of the same type, you can declare them like this:

```go
func plusPlus(a, b, c int) int {
    return a + b + c
}
```

## Multiple Return Values

Functions in Go can return multiple values. This is often used to return a result and an error value.

```go
func vals() (int, int) {
    return 3, 7
}
```

When calling a function with multiple return values, you can assign them to multiple variables.

```go
a, b := vals()
```

If you only need one of the return values, you can use the blank identifier `_` to discard the other.

```go
_, c := vals()
```

## Variadic Functions

A variadic function is a function that can accept a variable number of arguments.

```go
func sum(nums ...int) {
    // ...
}
```

You can call a variadic function with any number of arguments.

```go
sum(1, 2)
sum(1, 2, 3)
```

If you have a slice of values, you can pass them to a variadic function by using the `...` syntax.

```go
nums := []int{1, 2, 3, 4}
sum(nums...)
```

## Running the Code

To run the code in `4_Functions.go`, use the following command:

```bash
go run "4 Functions/4_Functions.go"
```

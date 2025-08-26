
# 1. Variables in Go

This section covers the basics of declaring and using variables in the Go programming language.

## What are Variables?

Variables are used to store data in a program. Each variable has a name, a type, and a value. The type of a variable determines what kind of data it can hold (e.g., numbers, text, etc.).

## Declaring Variables

There are a few ways to declare variables in Go:

### 1. Using the `var` keyword

You can declare a variable using the `var` keyword, followed by the variable name and its type.

```go
var age int
```

If you don't assign a value to the variable when you declare it, it will be given a "zero value" for its type. For example, the zero value for an `int` is `0`, and for a `string` it is an empty string `""`.

You can also initialize the variable with a value when you declare it:

```go
var name string = "Alice"
```

### 2. Type Inference

Go can often infer the type of a variable from the value you assign to it. This means you don't always have to explicitly specify the type.

```go
var city = "New York" // Go infers that 'city' is a string
```

### 3. Short Variable Declaration

The most common way to declare and initialize a variable in Go is by using the short variable declaration `:=`.

```go
country := "USA"
```

This syntax is concise and is preferred in most cases. However, it can only be used inside functions.

## Basic Data Types

Go has several built-in data types, including:

- `int`: Integers (whole numbers).
- `float64`: Floating-point numbers (numbers with a decimal point).
- `bool`: Boolean values (`true` or `false`).
- `string`: Text.

## Multiple Variable Declarations

You can declare multiple variables in a single statement:

```go
var a, b, c int = 1, 2, 3
```

Or, for variables of different types, you can use a block:

```go
var (
    firstName = "John"
    lastName  = "Doe"
    height    = 1.80
)
```

## Running the Code

To run the code in `1_Variables.go`, you can use the following command in your terminal:

```bash
go run "1 Variables/1_Variables.go"
```

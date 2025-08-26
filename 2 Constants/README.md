
# 2. Constants in Go

This section explains how to declare and use constants in Go.

## What are Constants?

Constants are similar to variables, but their values cannot be changed after they are declared. They are fixed values that are known at compile time.

## Declaring Constants

Constants are declared using the `const` keyword.

```go
const Pi = 3.14
```

Constants can be of any of the basic data types like `int`, `float64`, `bool`, or `string`.

```go
const World = "World"
const Truth = true
```

## Typed and Untyped Constants

Constants can be *typed* or *untyped*.

- **Typed Constants**: These are constants that are created with a specific type.

  ```go
  const typedInt int = 10
  ```

- **Untyped Constants**: These are constants that are not given a specific type. Their type is inferred from the context in which they are used.

  ```go
  const untypedInt = 20
  ```

Untyped constants are more flexible because they can be used in different contexts where different types are expected.

## Constant Expressions

Constants can be formed by performing arithmetic operations with other constants.

```go
const (
    Big   = 1 << 100
    Small = Big >> 99
)
```

## Running the Code

To run the code in `2_Constants.go`, you can use the following command in your terminal:

```bash
go run "2 Constants/2_Constants.go"
```

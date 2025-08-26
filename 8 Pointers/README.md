
# 8. Pointers in Go

This section explains the concept of pointers in Go.

## What are Pointers?

A pointer is a variable that stores the memory address of another variable. This allows you to indirectly access and modify the value of the original variable.

## Declaring Pointers

You can declare a pointer by using the `*` operator followed by the type of the variable that the pointer will point to.

```go
var p *int // Declares a pointer to an integer
```

## The `&` and `*` Operators

- The `&` operator is used to get the memory address of a variable.
- The `*` operator is used to dereference a pointer, which means accessing the value stored at the memory address that the pointer holds.

## Pass by Value vs. Pass by Reference

In Go, arguments are passed to functions by value. This means that the function receives a copy of the arguments, and any changes made to the arguments inside the function will not affect the original variables.

However, you can use pointers to achieve a "pass by reference" behavior. By passing a pointer to a function, the function can modify the value of the original variable.

### Example

Consider these two functions:

```go
func zeroval(ival int) {
    ival = 0
}

func zeroptr(iptr *int) {
    *iptr = 0
}
```

- `zeroval` takes an `int` as an argument. When you call this function, a copy of the integer is passed to it. Any changes to `ival` inside the function will not affect the original variable.

- `zeroptr` takes a pointer to an `int` as an argument. When you call this function with the memory address of a variable, the function can dereference the pointer and modify the original value.

```go
i := 1
fmt.Println("initial:", i)

zeroval(i)
fmt.Println("zeroval:", i) // i is still 1

zeroptr(&i)
fmt.Println("zeroptr:", i) // i is now 0
```

## Running the Code

To run the code in `8_Pointers.go`, use the following command:

```bash
go run "8 Pointers/8_Pointers.go"
```

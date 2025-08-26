# 3. Control Flow in Go

This section covers the control flow statements in Go, including `if/else`, `for`, and `switch`.

## `if/else` Statements

The `if/else` statement is used to execute a block of code if a certain condition is true, and another block of code if the condition is false.

```go
x := 10
if x > 5 {
    fmt.Println("x is greater than 5")
} else {
    fmt.Println("x is not greater than 5")
}
```

Go allows you to include an initialization statement in the `if` condition. The variable created in the initialization statement is only available within the scope of the `if/else` block.

```go
if y := 20; y > 10 {
    fmt.Println("y is greater than 10")
}
```

## `for` Loop

Go has only one looping construct, the `for` loop. It can be used in several ways:

### 1. Classic `for` loop

This is the traditional `for` loop with an initializer, a condition, and a post-statement.

```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

### 2. "While" loop

You can omit the initializer and post-statement to create a loop that behaves like a `while` loop in other languages.

```go
j := 0
for j < 5 {
    fmt.Println(j)
    j++
}
```

### 3. Infinite loop

By omitting the condition, you can create an infinite loop. You can use the `break` statement to exit the loop.

```go
for { 
    // This loop will run forever
    // Use 'break' to exit
}
```

## `switch` Statement

The `switch` statement is a more concise way to write a sequence of `if/else` statements.

```go
switch os := runtime.GOOS; os {
case "darwin":
    fmt.Println("OS X.")
case "linux":
    fmt.Println("Linux.")
default:
    fmt.Printf("%s.\n", os)
}
```

Key features of Go's `switch`:

- It automatically includes a `break` statement at the end of each case.
- You can have an initialization statement, similar to the `if` statement.
- You can use a `switch` with no condition to write a clean `if-then-else` chain.

```go
t := time.Now()
switch {
case t.Hour() < 12:
    fmt.Println("Good morning!")
case t.Hour() < 17:
    fmt.Println("Good afternoon.")
default:
    fmt.Println("Good evening.")
}
```

## Running the Code

To run the code in `3_Control_Flow.go`, use the following command:

```bash
go run "3 Control Flow/3_Control_Flow.go"
```



# 10. Goroutines in Go

This section introduces goroutines, which are a key feature of Go for concurrent programming.

## What are Goroutines?

A goroutine is a lightweight thread of execution. Goroutines are managed by the Go runtime, and they allow you to write concurrent programs that can perform multiple tasks at the same time.

## Creating a Goroutine

You can create a goroutine by using the `go` keyword followed by a function call.

```go
go f("goroutine")
```

This will start a new goroutine that executes the `f` function concurrently with the current goroutine.

You can also start a goroutine for an anonymous function.

```go
go func(msg string) {
    fmt.Println(msg)
}("going")
```

## Concurrency vs. Parallelism

- **Concurrency** is about dealing with multiple things at once.
- **Parallelism** is about doing multiple things at once.

Goroutines allow you to write concurrent code, and the Go runtime will automatically schedule them to run on one or more available CPU cores, achieving parallelism.

## Synchronization

When you start a goroutine, it runs independently of the main program. If the main program finishes before the goroutine, the goroutine will be terminated.

In the example code, `time.Sleep(time.Second)` is used to wait for the goroutines to finish. This is a simple way to wait, but for more complex scenarios, you should use synchronization mechanisms like channels or WaitGroups.

## Running the Code

To run the code in `10_Goroutines.go`, use the following command:

```bash
go run "10 Goroutines/10_Goroutines.go"
```

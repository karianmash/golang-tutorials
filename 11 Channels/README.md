
# 11. Channels in Go

This section covers channels, which are used for communication between goroutines.

## What are Channels?

Channels are the pipes that connect concurrent goroutines. You can send values into a channel from one goroutine and receive those values in another goroutine.

## Creating a Channel

You can create a channel using the `make` function.

```go
messages := make(chan string)
```

This creates a channel that can be used to send and receive values of type `string`.

## Sending and Receiving

- **Sending**: You can send a value into a channel using the `<-` operator.

  ```go
  messages <- "ping"
  ```

- **Receiving**: You can receive a value from a channel using the `<-` operator.

  ```go
  msg := <-messages
  ```

## Blocking

By default, sending and receiving on a channel are blocking operations.

- When you send a value to a channel, the sending goroutine will block until another goroutine is ready to receive the value.
- When you receive a value from a channel, the receiving goroutine will block until a value is sent to the channel.

This blocking behavior is a powerful synchronization mechanism. It allows you to coordinate the execution of goroutines without using explicit locks or condition variables.

## Buffered Channels

By default, channels are *unbuffered*, meaning that they will only accept a send if there is a corresponding receive ready to take the value. *Buffered* channels accept a limited number of values without a corresponding receiver for those values.

```go
ch := make(chan int, 100) // Creates a buffered channel with a capacity of 100
```

## Running the Code

To run the code in `11_Channels.go`, use the following command:

```bash
go run "11 Channels/11_Channels.go"
```

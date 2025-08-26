
# 9. Interfaces in Go

This section introduces interfaces, a powerful feature in Go for specifying behavior.

## What are Interfaces?

An interface is a collection of method signatures. A type is said to implement an interface if it has all the methods that are defined in the interface.

Interfaces provide a way to achieve polymorphism in Go. They allow you to write functions that can work with different types, as long as those types implement the same interface.

## Defining an Interface

You can define an interface using the `type` and `interface` keywords.

```go
type geometry interface {
    area() float64
    perim() float64
}
```

This defines an interface called `geometry` with two methods: `area()` and `perim()`, both of which return a `float64`.

## Implementing an Interface

A type implements an interface by implementing its methods. There is no explicit declaration of intent, no "implements" keyword.

```go
type rect struct {
    width, height float64
}

func (r rect) area() float64 {
    return r.width * r.height
}

func (r rect) perim() float64 {
    return 2*r.width + 2*r.height
}
```

In this example, the `rect` type implements the `geometry` interface because it has both the `area()` and `perim()` methods.

## Using Interfaces

Once a type implements an interface, you can use it as a value of that interface type.

```go
func measure(g geometry) {
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())
}
```

The `measure` function takes a `geometry` interface as an argument. This means you can pass any type that implements the `geometry` interface to this function.

```go
r := rect{width: 3, height: 4}
c := circle{radius: 5}

measure(r)
measure(c)
```

## Running the Code

To run the code in `9_Interfaces.go`, use the following command:

```bash
go run "9 Interfaces/9_Interfaces.go"
```

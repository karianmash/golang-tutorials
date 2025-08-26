
# 6. Maps in Go

This section covers maps, which are Go's built-in associative data type.

## What are Maps?

A map is a collection of key-value pairs. It is an unordered collection, and the keys must be unique.

## Creating a Map

You can create a map using the `make` function.

```go
m := make(map[string]int)
```

This creates a map where the keys are `string`s and the values are `int`s.

## Working with Maps

### Setting Key-Value Pairs

You can set a key-value pair using the following syntax:

```go
m["k1"] = 7
m["k2"] = 13
```

### Getting a Value

You can get the value for a key using the key as an index:

```go
v1 := m["k1"]
```

### `len`

The `len` function returns the number of key-value pairs in the map.

```go
fmt.Println("len:", len(m))
```

### `delete`

The `delete` function removes a key-value pair from the map.

```go
delete(m, "k2")
```

### Checking for Presence

When you get a value from a map, you can also get a second return value that indicates whether the key was present in the map.

```go
_, prs := m["k2"]
```

The second value (`prs`) will be `true` if the key exists, and `false` otherwise.

## Declaring and Initializing

You can declare and initialize a map in a single line.

```go
n := map[string]int{"foo": 1, "bar": 2}
```

## Running the Code

To run the code in `6_Maps.go`, use the following command:

```bash
go run "6 Maps/6_Maps.go"
```

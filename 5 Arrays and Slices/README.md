
# 5. Arrays and Slices in Go

This section covers arrays and slices, which are used to store collections of data.

## Arrays

An array is a fixed-size collection of elements of the same type. The size of the array is part of its type.

### Declaration

You can declare an array by specifying the type of its elements and its length.

```go
var a [5]int // Declares an array of 5 integers
```

By default, an array is zero-valued, meaning that all its elements are initialized to the zero value of their type (e.g., `0` for `int`, `""` for `string`).

### Initialization

You can initialize an array with values when you declare it.

```go
b := [5]int{1, 2, 3, 4, 5}
```

### Accessing Elements

You can access the elements of an array using an index.

```go
a[4] = 100
fmt.Println(a[4])
```

## Slices

A slice is a more flexible view into the elements of an array. Slices are much more common in Go than arrays because they can be resized.

### Creating a Slice

You can create a slice using the `make` function.

```go
s := make([]string, 3) // Creates a slice of 3 strings
```

### `append`

You can add elements to a slice using the `append` function.

```go
s = append(s, "d")
s = append(s, "e", "f")
```

### `copy`

You can copy the elements of one slice to another using the `copy` function.

```go
c := make([]string, len(s))
copy(c, s)
```

### Slicing a Slice

You can create a new slice from a portion of an existing slice using the `slice[low:high]` syntax.

```go
l := s[2:5] // Elements from index 2 up to (but not including) index 5
l = s[:5]  // Elements from the beginning up to index 5
l = s[2:]  // Elements from index 2 to the end
```

## Multi-dimensional Slices

You can create multi-dimensional data structures using slices. The inner slices can have different lengths.

```go
twoDSlice := make([][]int, 3)
for i := 0; i < 3; i++ {
    innerLen := i + 1
    twoDSlice[i] = make([]int, innerLen)
    // ...
}
```

## Running the Code

To run the code in `5_Arrays_and_Slices.go`, use the following command:

```bash
go run "5 Arrays and Slices/5_Arrays_and_Slices.go"
```

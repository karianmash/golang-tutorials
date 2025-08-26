package main

import "fmt"

func main() {
	// --- Arrays ---

	// In Go, an array is a numbered sequence of elements of a specific length.
	// The type of elements and length are both part of the array's type.

	// Here we create an array 'a' that will hold exactly 5 ints.
	// The type of elements and length are both part of the array's type.
	// By default an array is zero-valued, which for ints means 0s.
	var a [5]int
	fmt.Println("emp:", a) // Output: emp: [0 0 0 0 0]

	// We can set a value at an index using the `array[index] = value` syntax,
	// and get a value with `array[index]`.
	a[4] = 100
	fmt.Println("set:", a)    // Output: set: [0 0 0 0 100]
	fmt.Println("get:", a[4]) // Output: get: 100

	// The builtin `len` returns the length of an array.
	fmt.Println("len:", len(a)) // Output: len: 5

	// You can also declare and initialize an array in one line.
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b) // Output: dcl: [1 2 3 4 5]

	// Array types are one-dimensional, but you can compose types to build multi-dimensional data structures.
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD) // Output: 2d:  [[0 1 2] [1 2 3]]

	// --- Slices ---

	// A slice is a dynamically-sized, flexible view into the elements of an array.
	// In practice, slices are much more common than arrays.

	// To create an empty slice with non-zero length, use the builtin `make`.
	// Here we make a slice of strings of length 3 (initially zero-valued).
	s := make([]string, 3)
	fmt.Println("emp:", s) // Output: emp: [  ]

	// We can set and get just like with arrays.
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)    // Output: set: [a b c]
	fmt.Println("get:", s[2]) // Output: get: c

	// `len` returns the length of the slice as expected.
	fmt.Println("len:", len(s)) // Output: len: 3

	// In addition to these basic operations, slices support several more that make them richer than arrays.
	// One is the builtin `append`, which returns a slice containing one or more new values.
	// Note that we need to accept a return value from `append` as we may get a new slice value.
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s) // Output: apd: [a b c d e f]

	// Slices can also be `copy`'d. Here we create an empty slice `c` of the same length as `s` and copy into `c` from `s`.
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c) // Output: cpy: [a b c d e f]

	// Slices support a "slice" operator with the syntax `slice[low:high]`.
	// For example, this gets a slice of the elements `s[2]`, `s[3]`, and `s[4]`.
	l := s[2:5]
	fmt.Println("sl1:", l) // Output: sl1: [c d e]

	// This slices up to (but excluding) `s[5]`.
	l = s[:5]
	fmt.Println("sl2:", l) // Output: sl2: [a b c d e]

	// And this slices up from (and including) `s[2]`.
	l = s[2:]
	fmt.Println("sl3:", l) // Output: sl3: [c d e f]

	// We can declare and initialize a variable for a slice in a single line as well.
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t) // Output: dcl: [g h i]

	// Slices can be composed into multi-dimensional data structures.
	// The length of the inner slices can vary, unlike with multi-dimensional arrays.
	twoDSlice := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoDSlice[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoDSlice[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoDSlice) // Output: 2d:  [[0] [1 2] [2 3 4]]
}

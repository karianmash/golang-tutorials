
# 15. Testing in Go

This section covers how to write tests in Go.

## Writing a Test

In Go, tests are written in files that end with `_test.go`. The tests are written as functions that start with `Test` and take a `*testing.T` as a parameter.

```go
func TestIntMinBasic(t *testing.T) {
    ans := IntMin(2, -2)
    if ans != -2 {
        t.Errorf("IntMin(2, -2) = %d; want -2", ans)
    }
}
```

If the test fails, you can use the `t.Errorf` function to report an error.

## Table-Driven Tests

A common way to write tests in Go is to use a table-driven approach. This allows you to define a set of test cases and run them in a loop.

```go
func TestIntMinTableDriven(t *testing.T) {
    var tests = []struct {
        a, b int
        want int
    }{
        {0, 1, 0},
        {1, 0, 0},
        {2, -2, -2},
        {0, -1, -1},
        {-1, 0, -1},
    }

    for _, tt := range tests {
        t.Run("", func(t *testing.T) {
            ans := IntMin(tt.a, tt.b)
            if ans != tt.want {
                t.Errorf("got %d, want %d", ans, tt.want)
            }
        })
    }
}
```

## Running Tests

To run the tests in a package, you can use the `go test` command.

```bash
go test -v
```

The `-v` flag prints the name and status of each test.

## Running the Code

To run the tests in `15_Testing_test.go`, use the following command in the `15 Testing` directory:

```bash
cd "15 Testing"
go test -v
```

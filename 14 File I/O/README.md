# 14. File I/O in Go

This section covers how to read from and write to files in Go.

## Reading Files

### Reading an Entire File

The easiest way to read an entire file into memory is to use the `ioutil.ReadFile` function.

```go
dat, err := ioutil.ReadFile("dat")
```

### Finer-Grained Reading

For more control over reading, you can open a file using `os.Open`.

```go
f, err := os.Open("dat")
```

This returns an `os.File` object, which you can use to read the file in smaller chunks.

```go
b1 := make([]byte, 5)
n1, err := f.Read(b1)
```

You can also seek to a specific position in the file using `f.Seek`.

### Buffered Reading

The `bufio` package provides a buffered reader, which can be more efficient for many small reads.

```go
r4 := bufio.NewReader(f)
b4, err := r4.Peek(5)
```

## Writing Files

### Writing an Entire File

You can write a slice of bytes to a file using `ioutil.WriteFile`.

```go
d1 := []byte("hello\ngo\n")
err = ioutil.WriteFile("dat1", d1, 0644)
```

### Finer-Grained Writing

For more control over writing, you can create a file using `os.Create`.

```go
f, err := os.Create("dat2")
```

This returns an `os.File` object that you can use to write to the file.

```go
n2, err = f.Write(d2)
n3, err = f.WriteString("writes\n")
```

### Buffered Writing

The `bufio` package also provides a buffered writer.

```go
w := bufio.NewWriter(f)
n4, err := w.WriteString("buffered\n")
w.Flush() // Don't forget to flush the buffer
```

## Error Handling

It is important to check for errors when working with files. The `check` helper function in the example shows a simple way to handle errors.

## Closing Files

It is important to close files when you are finished with them. The `defer` statement is often used to ensure that a file is closed, even if an error occurs.

```go
f, err := os.Open("dat")
defer f.Close()
```

## Running the Code

To run the code in `14_File_IO.go`, use the following command:

```bash
go run "14 File I/O/14_File_IO.go"
```

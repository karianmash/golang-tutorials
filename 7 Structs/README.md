
# 7. Structs in Go

This section covers structs, which are used to create custom data types in Go.

## What are Structs?

A struct is a collection of fields. It is a way to group together related data to form a record.

## Defining a Struct

You can define a struct using the `type` and `struct` keywords.

```go
type person struct {
    name string
    age  int
}
```

This defines a new type called `person` with two fields: `name` of type `string` and `age` of type `int`.

## Creating a Struct

You can create an instance of a struct in several ways:

```go
// 1. Using the struct literal syntax
fmt.Println(person{"Bob", 20})

// 2. Naming the fields
fmt.Println(person{name: "Alice", age: 30})

// 3. Omitted fields will be zero-valued
fmt.Println(person{name: "Fred"})

// 4. Using the `&` prefix to get a pointer to the struct
fmt.Println(&person{name: "Ann", age: 40})
```

## Accessing Fields

You can access the fields of a struct using the dot `.` operator.

```go
s := person{name: "Sean", age: 50}
fmt.Println(s.name) // Access the name field
```

If you have a pointer to a struct, you can still use the dot operator to access its fields. The pointer will be automatically dereferenced.

```go
sp := &s
fmt.Println(sp.age) // Access the age field through a pointer
```

## Mutability

Structs are mutable, which means you can change the values of their fields.

```go
sp.age = 51
```

## Running the Code

To run the code in `7_Structs.go`, use the following command:

```bash
go run "7 Structs/7_Structs.go"
```

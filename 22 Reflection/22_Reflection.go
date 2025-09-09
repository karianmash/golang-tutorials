package main

import (
	"fmt"
	"reflect"
)

// Person struct for demonstration
type Person struct {
	Name string
	Age  int
}

func main() {
	fmt.Println("--- Inspecting Types and Values ---")
	var x float64 = 3.14159
	v := reflect.ValueOf(x)
	t := reflect.TypeOf(x)

	fmt.Printf("Value: %v, Type: %v, Kind: %v\n", v, t, v.Kind())

	// Getting underlying value
	fmt.Printf("Underlying value (float64): %f\n", v.Float())

	fmt.Println("\n--- Inspecting Structs ---")
	p := Person{"Alice", 30}
	pVal := reflect.ValueOf(p)
	pType := reflect.TypeOf(p)

	fmt.Printf("Struct Value: %v, Type: %v, Kind: %v\n", pVal, pType, pVal.Kind())

	// Iterate over struct fields
	for i := 0; i < pVal.NumField(); i++ {
		field := pVal.Field(i)
		fieldType := pType.Field(i)
		fmt.Printf("Field %d: Name=%s, Type=%v, Kind=%v, Value=%v\n",
			i, fieldType.Name, fieldType.Type, field.Kind(), field.Interface())
	}

	// Accessing a specific field by name
	nameField := pVal.FieldByName("Name")
	if nameField.IsValid() {
		fmt.Printf("Name field value: %v\n", nameField.Interface())
	}

	fmt.Println("\n--- Modifying Values (Requires Addressability) ---")
	// To modify a value using reflection, you need to pass a pointer to it.
	y := 10
	yPtr := reflect.ValueOf(&y) // Get ValueOf the pointer
	yVal := yPtr.Elem()         // Get the element the pointer points to

	fmt.Printf("Original y: %d, yVal.CanSet(): %t\n", y, yVal.CanSet())

	if yVal.CanSet() {
		yVal.SetInt(20)
		fmt.Printf("Modified y: %d\n", y)
	}

	// Modifying a struct field
	fmt.Println("\n--- Modifying Struct Fields (Requires Addressability) ---")
	personPtr := reflect.ValueOf(&p) // Get ValueOf the pointer to the struct
	personVal := personPtr.Elem()    // Get the element the pointer points to

	fmt.Printf("Original Person: %+v\n", p)

	nameFieldMod := personVal.FieldByName("Name")
	ageFieldMod := personVal.FieldByName("Age")

	if nameFieldMod.IsValid() && nameFieldMod.CanSet() {
		nameFieldMod.SetString("Bob")
	}
	if ageFieldMod.IsValid() && ageFieldMod.CanSet() {
		ageFieldMod.SetInt(35)
	}

	fmt.Printf("Modified Person: %+v\n", p)

	fmt.Println("\n--- Calling Methods (Requires Addressability) ---")
	// Define a method for Person
	type Greeter struct {
		Name string
	}

	func(g Greeter) SayHello()
	{
		fmt.Printf("Hello, my name is %s!\n", g.Name)
	}

	g := Greeter{"Charlie"}
	gVal := reflect.ValueOf(g)
	method := gVal.MethodByName("SayHello")

	if method.IsValid() {
		method.Call(nil) // Call method with no arguments
	}

	// Calling a method that takes arguments
	func(g Greeter) Greet(message
	string) {
		fmt.Printf("Hello %s, %s!\n", g.Name, message)
	}

	gValWithArgs := reflect.ValueOf(g)
	methodWithArgs := gValWithArgs.MethodByName("Greet")

	if methodWithArgs.IsValid() {
		args := []reflect.Value{reflect.ValueOf("nice to meet you")}
		methodWithArgs.Call(args)
	}
}

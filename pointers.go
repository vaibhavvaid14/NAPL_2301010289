package main

import "fmt"

// Define a struct for part (3)
type Person struct {
	Name string
	Age  int
}

// Function accepting a pointer parameter to modify the original variable (pass-by-reference)
func increment(val *int) {
	*val = *val + 10
}

func main() {
	// -------------------------------------------------------------
	// (1) Declare a variable, print its address using & and value using *
	// -------------------------------------------------------------
	fmt.Println("--- (1) Basic Pointer Operations ---")
	num := 42
	ptr := &num // ptr holds the memory address of num

	fmt.Printf("Value of num: %d\n", num)
	fmt.Printf("Address of num (&num): %p\n", &num)
	fmt.Printf("Address stored in ptr: %p\n", ptr)
	fmt.Printf("Value accessed via dereferencing (*ptr): %d\n\n", *ptr)

	// -------------------------------------------------------------
	// (2) Pass-by-reference using a pointer parameter
	// -------------------------------------------------------------
	fmt.Println("--- (2) Pass-by-Reference Function ---")
	count := 5
	fmt.Printf("Value before function call: %d\n", count)

	increment(&count) // Pass memory address of count

	fmt.Printf("Value after function call: %d\n\n", count)

	// -------------------------------------------------------------
	// (3) Allocate a struct using new() and modify fields
	// -------------------------------------------------------------
	fmt.Println("--- (3) Struct Allocation with new() ---")
	// new(Person) allocates memory for Person and returns a pointer (*Person)
	p := new(Person)

	// Access and modify fields through the pointer
	// Go automatically dereferences pointers to structs (e.g., p.Name is equivalent to (*p).Name)
	p.Name = "Alice"
	p.Age = 25

	fmt.Printf("Struct pointer address: %p\n", p)
	fmt.Printf("Struct value (*p): %+v\n", *p)
	fmt.Printf("Modified Fields -> Name: %s, Age: %d\n", p.Name, p.Age)
}
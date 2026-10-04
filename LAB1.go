package main

import (
	"fmt"
	"runtime"
)

func main() {
	// Print Go runtime version (helps verify installation when running `go run`).
	fmt.Println("Go runtime version:", runtime.Version())

	// Integer operations
	a, b := 12, 5
	fmt.Printf("Integers: a=%d, b=%d\n", a, b)
	fmt.Printf("Int add: %d + %d = %d\n", a, b, a+b)
	fmt.Printf("Int sub: %d - %d = %d\n", a, b, a-b)
	fmt.Printf("Int mul: %d * %d = %d\n", a, b, a*b)

	// Floating-point operations
	x, y := 3.5, 1.2
	fmt.Printf("Floats: x=%.2f, y=%.2f\n", x, y)
	fmt.Printf("Float add: %.2f + %.2f = %.2f\n", x, y, x+y)
	fmt.Printf("Float sub: %.2f - %.2f = %.2f\n", x, y, x-y)
	fmt.Printf("Float mul: %.2f * %.2f = %.2f\n", x, y, x*y)
}


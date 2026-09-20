package main

import (
	"fmt"
	"go-custom-pkg/myutils"
)

func main() {
	fmt.Println("=== Custom Reusable Package Demo ===")

	// String Functions Demonstration
	text := "Hello World"
	fmt.Printf("\n--- String Operations ---\n")
	fmt.Printf("Original String: %s\n", text)
	fmt.Printf("Reversed String: %s\n", myutils.Reverse(text))
	fmt.Printf("Vowel Count:     %d\n", myutils.CountVowels(text))

	// Math Functions Demonstration
	fmt.Printf("\n--- Math Operations ---\n")
	num := 5
	fact, err := myutils.Factorial(num)
	if err != nil {
		fmt.Printf("Error calculating factorial: %v\n", err)
	} else {
		fmt.Printf("Factorial of %d: %d\n", num, fact)
	}

	base, exp := 2, 8
	fmt.Printf("%d raised to power %d: %d\n", base, exp, myutils.Power(base, exp))
}
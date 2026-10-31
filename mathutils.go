package main

import (
    "fmt"
    "math"
    "strings"
)

// 1. Division scenario: handles division by zero
func Divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("division error: cannot divide %f by zero", a)
    }
    return a / b, nil
}

// 2. Square root scenario: handles negative numbers
func ComputeSqrt(x float64) (float64, error) {
    if x < 0 {
        return 0, fmt.Errorf("math error: cannot compute square root of negative number %f", x)
    }
    return math.Sqrt(x), nil
}

// 3. Input validation scenario: handles empty or blank strings
func ProcessInput(input string) (string, error) {
    trimmed := strings.TrimSpace(input)
    if trimmed == "" {
        return "", fmt.Errorf("validation error: input string cannot be empty or whitespace")
    }
    return fmt.Sprintf("Processed: %s", trimmed), nil
}

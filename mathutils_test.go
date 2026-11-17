package main

import (
    "testing"
)

// Test for Division (Valid + Error Case)
func TestDivide(t *testing.T) {
    // Valid test case
    val, err := Divide(10, 2)
    if err != nil || val != 5 {
        t.Errorf("Expected 5, nil; got %f, %v", val, err)
    }

    // Error case: Division by zero
    _, err = Divide(10, 0)
    if err == nil {
        t.Error("Expected division by zero error, got nil")
    }
}

// Test for Square Root (Valid + Error Case)
func TestComputeSqrt(t *testing.T) {
    // Valid test case
    val, err := ComputeSqrt(9)
    if err != nil || val != 3 {
        t.Errorf("Expected 3, nil; got %f, %v", val, err)
    }

    // Error case: Negative number
    _, err = ComputeSqrt(-4)
    if err == nil {
        t.Error("Expected negative square root error, got nil")
    }
}

// Test for Empty Input (Valid + Error Case)
func TestProcessInput(t *testing.T) {
    // Valid test case
    res, err := ProcessInput("  Hello Go  ")
    if err != nil || res != "Processed: Hello Go" {
        t.Errorf("Expected 'Processed: Hello Go', nil; got %s, %v", res, err)
    }

    // Error case: Empty string
    _, err = ProcessInput("   ")
    if err == nil {
        t.Error("Expected empty input error, got nil")
    }
}

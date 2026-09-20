package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	minValue = -1000000000
	maxValue = 1000000000
)

func readLine(scanner *bufio.Scanner, prompt string) (string, bool) {
	fmt.Print(prompt)
	if !scanner.Scan() {
		return "", false
	}
	return strings.TrimSpace(scanner.Text()), true
}

func readChoice(scanner *bufio.Scanner) (int, bool) {
	for {
		value, ok := readLine(scanner, "Choose a type (1 = integer, 2 = floating-point, 0 = exit): ")
		if !ok {
			return 0, false
		}

		choice, err := strconv.Atoi(value)
		if err == nil && choice >= 0 && choice <= 2 {
			return choice, true
		}
		fmt.Println("Invalid choice. Enter 0, 1, or 2.")
	}
}

func readInteger(scanner *bufio.Scanner, name string) (int, bool) {
	for {
		value, ok := readLine(scanner, "Enter "+name+" (between -1000000000 and 1000000000): ")
		if !ok {
			return 0, false
		}

		number, err := strconv.Atoi(value)
		if err == nil && number >= minValue && number <= maxValue {
			return number, true
		}
		fmt.Println("Invalid integer. Enter a numeric value within the allowed range.")
	}
}

func readFloat(scanner *bufio.Scanner, name string) (float64, bool) {
	for {
		value, ok := readLine(scanner, "Enter "+name+" (between -1000000000 and 1000000000): ")
		if !ok {
			return 0, false
		}

		number, err := strconv.ParseFloat(value, 64)
		if err == nil && !math.IsNaN(number) && !math.IsInf(number, 0) && number >= minValue && number <= maxValue {
			return number, true
		}
		fmt.Println("Invalid floating-point value. Enter a finite numeric value within the allowed range.")
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		choice, ok := readChoice(scanner)
		if !ok || choice == 0 {
			fmt.Println("Calculator closed.")
			return
		}

		if choice == 1 {
			first, ok := readInteger(scanner, "the first integer")
			if !ok {
				return
			}
			second, ok := readInteger(scanner, "the second integer")
			if !ok {
				return
			}

			fmt.Printf("Addition: %d + %d = %d\n", first, second, first+second)
			fmt.Printf("Subtraction: %d - %d = %d\n", first, second, first-second)
			fmt.Printf("Multiplication: %d * %d = %d\n", first, second, first*second)
		} else {
			first, ok := readFloat(scanner, "the first floating-point number")
			if !ok {
				return
			}
			second, ok := readFloat(scanner, "the second floating-point number")
			if !ok {
				return
			}

			fmt.Printf("Addition: %.2f + %.2f = %.2f\n", first, second, first+second)
			fmt.Printf("Subtraction: %.2f - %.2f = %.2f\n", first, second, first-second)
			fmt.Printf("Multiplication: %.2f * %.2f = %.2f\n", first, second, first*second)
		}

		fmt.Println()
	}
}

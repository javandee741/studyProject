package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func add(a, b float64) float64 {
	return a + b
}

func subtract(a, b float64) float64 {
	return a - b
}

func multiply(a, b float64) float64 {
	return a * b
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Input format: <number1> <operation> <number2>")
	fmt.Println("Example: 5 + 3")
	fmt.Print("Enter expression: ")

	scanner.Scan()
	input := scanner.Text()
	parts := strings.Fields(input)

	if len(parts) < 3 {
		fmt.Println("Error: too few arguments. Please use the format: <number> <operator> <number>")
		return
	}
	if len(parts) > 3 {
		fmt.Println("Error: too many arguments. Please use the format: <number> <operator> <number>")
		return
	}

	a, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		fmt.Printf("Error parsing first number: %v\n", err)
		return
	}

	op := parts[1]

	b, err := strconv.ParseFloat(parts[2], 64)
	if err != nil {
		fmt.Printf("Error parsing second number: %v\n", err)
		return
	}

	var result float64

	switch op {
	case "+":
		result = add(a, b)
	case "-":
		result = subtract(a, b)
	case "*":
		result = multiply(a, b)
	case "/":
		result, err = divide(a, b)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
	default:
		fmt.Printf("Unknown operation: %s\n", op)
		return
	}

	fmt.Printf("Result: %f\n", result)
}

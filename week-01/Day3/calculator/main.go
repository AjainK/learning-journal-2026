package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("=== Simple Calculator ===")
	fmt.Print("Enter first number: ")
	input1, _ := reader.ReadString('\n')
	input1 = strings.TrimSpace(input1)
	num1, err := strconv.ParseFloat(input1, 64)
	if err != nil {
		fmt.Println("Please enter a valid number.")
		return
	}

	fmt.Print("Enter second number: ")
	input2, _ := reader.ReadString('\n')
	input2 = strings.TrimSpace(input2)
	num2, err := strconv.ParseFloat(input2, 64)
	if err != nil {
		fmt.Println("Please enter a valid number.")
		return
	}

	fmt.Print("Enter an operator (+, -, *, /): ")
	operator, _ := reader.ReadString('\n')
	operator = strings.TrimSpace(operator)
	result, err := calculate(num1, num2, operator)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Result: %.2f\n", result)

}

func calculate(num1, num2 float64, operator string) (float64, error) {
	switch operator {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		if num2 == 0 {
			return 0, fmt.Errorf("cannot divide by zero")
		}
		return num1 / num2, nil
	default:
		return 0, fmt.Errorf("invalid operator")
	}
}

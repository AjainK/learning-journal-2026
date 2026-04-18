package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isEvenOrOdd(n int) string {
	if n%2 == 0 {
		return "Even"
	} else {
		return "Odd"
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("=== Even or Odd Checker ===")
	fmt.Print("Enter a number: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	input = strings.TrimSpace(input)
	number, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Please enter a valid number.")
		return
	}
	status := isEvenOrOdd(number)
	fmt.Printf("The given number %d is %s\n", number, status)
}

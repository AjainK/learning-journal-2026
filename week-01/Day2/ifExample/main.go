package main

import "fmt"

func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	if sum != 0 {
		fmt.Println(sum)
	} else {
		fmt.Println("Sum is zero")
	}

}

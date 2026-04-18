package main

import "fmt"

func main() {
	const boilingF, freezingF = 212.0, 32.0
	fmt.Printf("%g°F = %g°C \n", boilingF, ftoc(boilingF))
	fmt.Printf("%g°F = %g°C \n", freezingF, ftoc(freezingF))

}

func ftoc(f float64) float64 {
	return (f - 32.0) * 5.0 / 9.0
}

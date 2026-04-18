// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"math/rand"
// 	"os"
// 	"strconv"
// 	"strings"
// 	"time"
// )

// func main() {
// 	source := rand.NewSource(time.Now().UnixNano())
// 	r := rand.New(source)
// 	secretNumber := r.Intn(100) + 1

// 	fmt.Println("Welcome to the Number Guessing Game!")
// 	fmt.Println("I'm thinking of a number between 1 and 100.")
// 	fmt.Println("Can you guess it?")

// 	reader := bufio.NewReader(os.Stdin)
// 	attempts := 0

// 	for {
// 		fmt.Print("\nEnter your guess: ")
// 		input, _ := reader.ReadString('\n')
// 		input = strings.TrimSpace(input)

// 		guess, err := strconv.Atoi(input)
// 		if err != nil {
// 			fmt.Println("Please enter a valid number.")
// 			continue
// 		}

// 		if guess < 1 || guess > 100 {
// 			fmt.Println("Please enter a number between 1 and 100.")
// 			continue
// 		}

// 		attempts++

// 		if guess == secretNumber {
// 			fmt.Printf("🎉 Correct! You guessed it in %d attempt(s)!\n", attempts)
// 			break
// 		} else if guess < secretNumber {
// 			fmt.Println("📈 Too low! Try higher.")
// 		} else {
// 			fmt.Println("📉 Too high! Try lower.")
// 		}
// 	}
// }

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// ── FizzBuzz ──────────────────────────────────────────────────────────────────

func fizzBuzz(n int) string {
	switch {
	case n%15 == 0:
		return "FizzBuzz"
	case n%3 == 0:
		return "Fizz"
	case n%5 == 0:
		return "Buzz"
	default:
		return strconv.Itoa(n)
	}
}

func runFizzBuzz() {
	fmt.Println("=== FizzBuzz (1–100) ===")
	for i := 1; i <= 100; i++ {
		fmt.Printf("%3s  ", fizzBuzz(i))
		if i%10 == 0 {
			fmt.Println()
		}
	}
	fmt.Println()
}

// ── Number guessing game ──────────────────────────────────────────────────────

func runGuessingGame() {
	fmt.Println("=== Number Guessing Game ===")
	fmt.Println("I've picked a number between 1 and 100. Can you guess it?")
	fmt.Println()

	// seed random with current time so it's different every run
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	secret := rng.Intn(100) + 1 // 1–100 inclusive

	scanner := bufio.NewScanner(os.Stdin)
	attempts := 0
	maxAttempts := 7

	for attempts < maxAttempts {
		remaining := maxAttempts - attempts
		fmt.Printf("Guess (%d attempt(s) left): ", remaining)

		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("  Please enter a valid number.")
			continue
		}

		if guess < 1 || guess > 100 {
			fmt.Println("  Number must be between 1 and 100.")
			continue
		}

		attempts++

		switch {
		case guess == secret:
			fmt.Printf("\n  Correct! You got it in %d attempt(s).\n", attempts)
			return
		case guess < secret:
			fmt.Println("  Too low — go higher.")
		default:
			fmt.Println("  Too high — go lower.")
		}
	}

	fmt.Printf("\n  Out of attempts! The number was %d.\n", secret)
}

// ── Main ──────────────────────────────────────────────────────────────────────

func main() {
	runFizzBuzz()
	runGuessingGame()
}

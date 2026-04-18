package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name:")
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	name = strings.TrimSpace(name)
	now := time.Now()
	year, month, day := now.Date()
	location := now.Location()
	lastDayOfMonth := time.Date(year, month+1, 0, 0, 0, 0, 0, location).Day()
	daysLeft := lastDayOfMonth - day
	fmt.Printf("Hello, %s! There are %d days left in this month.\n", name, daysLeft)
}

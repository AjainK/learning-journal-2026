package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Task represents a to-do item
type Task struct {
	ID        int
	Title     string
	Completed bool
}

var tasks []Task
var nextID = 1

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("=== To-Do CLI ===")
	fmt.Println("Commands: add, list, mark-done, exit")
	fmt.Println()

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			continue
		}

		parts := strings.Fields(input) //["add", "Wake", "up", "early"]
		command := parts[0]
		switch command {
		case "add":
			if len(parts) < 2 {
				fmt.Println("Usage: add <task description>")
				continue
			}
			title := strings.Join(parts[1:], " ")
			addTask(title)

		case "list":
			listTasks()

		case "mark-done":
			if len(parts) < 2 {
				fmt.Println("Usage: mark-done <task id>")
				continue
			}
			id, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println("Invalid task ID")
				continue
			}
			markDone(id)

		case "exit":
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Unknown command. Use: add, list, mark-done, exit")
		}
	}
}

func addTask(title string) {
	task := Task{
		ID:        nextID,
		Title:     title,
		Completed: false,
	}
	tasks = append(tasks, task)
	fmt.Printf("✓ Task added with ID %d\n", nextID)
	nextID++
}

func listTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks yet!")
		return
	}

	fmt.Println("\n--- Tasks ---")
	for _, task := range tasks {
		status := "[ ]"
		if task.Completed {
			status = "[✓]"
		}
		fmt.Printf("%s %d. %s\n", status, task.ID, task.Title)
	}
	fmt.Println()
}

func markDone(id int) {
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Completed = true
			fmt.Printf("✓ Task %d marked as done\n", id)
			return
		}
	}
	fmt.Printf("Task %d not found\n", id)
}

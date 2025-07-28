package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/neel07sanghvi/todo-cli/internal/todo"
)

func main() {
	todoManager := todo.NewTodoManager()
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("=== Welcome to Todo CLI ===")
	fmt.Println("Commands: add, list, update, delete, complete, incomplete, help, exit")

	for {
		fmt.Print("\n> ")

		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())

		if input == "" {
			continue
		}

		parts := strings.SplitN(input, " ", 2)
		command := strings.ToLower(parts[0])

		switch command {
		case "add":
			if len(parts) < 2 {
				fmt.Println("Usage: add <task description>")
				continue
			}

			task := parts[1]
			id := todoManager.AddTodo(task)

			fmt.Printf("Todo added with ID: %d\n", id)

		case "list":
			todoManager.ListTodos()

		case "update":
			if len(parts) < 2 {
				fmt.Println("Usage: update <id> <new description>")
				continue
			}

			updateParts := strings.SplitN(parts[1], " ", 2)

			if len(updateParts) < 2 {
				fmt.Println("Usage: update <id> <new description>")
				continue
			}

			id, err := strconv.Atoi(updateParts[0])

			if err != nil {
				fmt.Println("Invalid ID. Please provide a valid number.")
				continue
			}

			newTask := updateParts[1]

			if todoManager.UpdateTodo(id, newTask) {
				fmt.Printf("Todo with ID %d updated successfully\n", id)
			} else {
				fmt.Printf("Todo with ID %d not found\n", id)
			}

		case "delete":
			if len(parts) < 2 {
				fmt.Println("Usage: delete <id>")
				continue
			}

			id, err := strconv.Atoi(parts[1])

			if err != nil {
				fmt.Println("Invalid ID. Please provide a valid number.")
				continue
			}

			if todoManager.DeleteTodo(id) {
				fmt.Printf("Todo with ID %d deleted successfully\n", id)
			} else {
				fmt.Printf("Todo with ID %d not found\n", id)
			}

		case "complete":
			if len(parts) < 2 {
				fmt.Println("Usage: complete <id>")
				continue
			}

			id, err := strconv.Atoi(parts[1])

			if err != nil {
				fmt.Println("Invalid ID. Please provide a valid number.")
				continue
			}

			if todoManager.CompleteTodo(id) {
				fmt.Printf("Todo with ID %d marked as completed\n", id)
			} else {
				fmt.Printf("Todo with ID %d not found\n", id)
			}

		case "incomplete":
			if len(parts) < 2 {
				fmt.Println("Usage: incomplete <id>")
				continue
			}

			id, err := strconv.Atoi(parts[1])

			if err != nil {
				fmt.Println("Invalid ID. Please provide a valid number.")
				continue
			}

			if todoManager.IncompleteTodo(id) {
				fmt.Printf("Todo with ID %d marked as incomplete\n", id)
			} else {
				fmt.Printf("Todo with ID %d not found\n", id)
			}

		case "help":
			printHelp()

		case "exit", "quit":
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Printf("Unknown command: %s. Type 'help' for available commands.\n", command)
		}
	}
}

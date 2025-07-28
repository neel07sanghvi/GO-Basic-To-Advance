package main

import "fmt"

func printHelp() {
	fmt.Println("\n=== Todo CLI Help ===")
	fmt.Println("Available commands:")
	fmt.Println("  add <task>       - Add a new todo item")
	fmt.Println("  list             - List all todo items")
	fmt.Println("  update <id> <task> - Update an existing todo item")
	fmt.Println("  delete <id>      - Delete a todo item")
	fmt.Println("  complete <id>    - Mark a todo item as completed")
	fmt.Println("  incomplete <id>  - Mark a todo item as incomplete")
	fmt.Println("  help             - Show this help message")
	fmt.Println("  exit/quit        - Exit the application")
	fmt.Println("\nExamples:")
	fmt.Println("  add Buy groceries")
	fmt.Println("  list")
	fmt.Println("  update 1 Buy groceries and cook dinner")
	fmt.Println("  complete 1")
	fmt.Println("  incomplete 1")
	fmt.Println("  delete 1")
}

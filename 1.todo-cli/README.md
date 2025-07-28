# Todo CLI Application

A simple command-line todo list application written in Go with CRUD operations.

## Project Structure

```
todo-cli/
├── main.go                 # Main application entry point
├── help.go                 # Help command implementation
├── go.mod                  # Go module file
├── README.md              # This file
└── internal/
    └── todo/
        ├── model.go       # Todo data structure
        └── manager.go     # Todo management logic
```

## Features

- **Create**: Add new todo items
- **Read**: List all todos with status and timestamps
- **Update**: Modify existing todo descriptions
- **Delete**: Remove todos from the list
- **Complete**: Mark todos as completed/incomplete
- **Statistics**: View completion statistics

## Installation

1. Clone or create the project:
```bash
mkdir todo-cli
cd todo-cli
```

2. Initialize the Go module:
```bash
go mod init github.com/yourusername/todo-cli
```

3. Create the folder structure and add the files as shown above.

4. Run the application:
```bash
go run .
```

## Usage

### Available Commands

- `add <task>` - Add a new todo item
- `list` - List all todo items
- `update <id> <new_task>` - Update an existing todo item
- `delete <id>` - Delete a todo item
- `complete <id>` - Mark a todo item as completed
- `incomplete <id>` - Mark a todo item as incomplete
- `help` - Show help message
- `exit` or `quit` - Exit the application

### Examples

```bash
> add Buy groceries
Todo added with ID: 1

> add Call dentist
Todo added with ID: 2

> list
=== Your Todos ===
1. [ ] Buy groceries (created: 2024-01-15 10:30)
2. [ ] Call dentist (created: 2024-01-15 10:31)

Total: 2 | Completed: 0 | Remaining: 2

> complete 1
Todo with ID 1 marked as completed

> incomplete 1
Todo with ID 1 marked as incomplete

> update 2 Schedule dentist appointment
Todo with ID 2 updated successfully

> list
=== Your Todos ===
1. [✓] Buy groceries (created: 2024-01-15 10:30) (completed: 2024-01-15 10:32)
2. [ ] Schedule dentist appointment (created: 2024-01-15 10:31)

Total: 2 | Completed: 1 | Remaining: 1

> delete 1
Todo with ID 1 deleted successfully
```

## Architecture

The application is structured with separation of concerns:

- **main.go**: Handles user input and CLI interface
- **help.go**: Contains help functionality
- **internal/todo/model.go**: Defines the Todo data structure
- **internal/todo/manager.go**: Implements business logic for managing todos

### Key Components

1. **Todo Model**: Represents individual todo items with ID, task, completion status, and timestamps
2. **TodoManager**: Manages the collection of todos with CRUD operations
3. **CLI Interface**: Interactive command-line interface for user interaction

## Data Storage

Currently, all data is stored in memory using Go maps and slices. Data will be lost when the application exits. This is intentional for the basic version - no database or file persistence is implemented.

## Future Enhancements

- File-based persistence (JSON/YAML)
- Database integration
- Priority levels for todos
- Due dates and reminders
- Categories and tags
- Search and filter functionality
- Export/import features
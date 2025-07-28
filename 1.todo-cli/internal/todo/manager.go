package todo

import (
	"fmt"
	"sort"
	"time"
)

// TodoManager manages the collection of todos
type TodoManager struct {
	todos  map[int]*Todo
	nextID int
}

// NewTodoManager creates a new TodoManager instance
func NewTodoManager() *TodoManager {
	return &TodoManager{
		todos:  make(map[int]*Todo),
		nextID: 1,
	}
}

// AddTodo adds a new todo item and returns its ID
func (tm *TodoManager) AddTodo(task string) int {
	todo := &Todo{
		ID:        tm.nextID,
		Task:      task,
		Completed: false,
		CreatedAt: time.Now(),
	}

	tm.todos[todo.ID] = todo
	tm.nextID++

	return todo.ID
}

// ListTodos displays all todo items
func (tm *TodoManager) ListTodos() {
	if len(tm.todos) == 0 {
		fmt.Println("No todos found. Add some todos to get started!")
		return
	}

	var ids []int
	for id := range tm.todos {
		ids = append(ids, id)
	}

	sort.Ints(ids)

	fmt.Println("\n=== Your Todos ===")

	for _, id := range ids {
		fmt.Println(tm.todos[id].String())
	}

	completed := tm.GetCompletedCount()
	total := len(tm.todos)

	fmt.Printf("\nTotal: %d | Completed: %d | Remaining: %d\n", total, completed, total-completed)
}

// UpdateTodo updates an existing todo item
func (tm *TodoManager) UpdateTodo(id int, newTask string) bool {
	todo, exists := tm.todos[id]

	if !exists {
		return false
	}

	todo.Task = newTask

	return true
}

// DeleteTodo removes a todo item
func (tm *TodoManager) DeleteTodo(id int) bool {
	_, exists := tm.todos[id]

	if !exists {
		return false
	}

	delete(tm.todos, id)

	return true
}

// CompleteTodo marks a todo as completed
func (tm *TodoManager) CompleteTodo(id int) bool {
	todo, exists := tm.todos[id]
	if !exists {
		return false
	}

	todo.MarkCompleted()

	return true
}

// IncompleteTodo marks a todo as incomplete
func (tm *TodoManager) IncompleteTodo(id int) bool {
	todo, exists := tm.todos[id]
	if !exists {
		return false
	}

	todo.MarkIncomplete()
	return true
}

// GetTodo retrieves a specific todo by ID
func (tm *TodoManager) GetTodo(id int) (*Todo, bool) {
	todo, exists := tm.todos[id]
	return todo, exists
}

// GetAllTodos returns all todos as a slice
func (tm *TodoManager) GetAllTodos() []*Todo {
	var todos []*Todo

	for _, todo := range tm.todos {
		todos = append(todos, todo)
	}

	// Sort by ID
	sort.Slice(todos, func(i, j int) bool {
		return todos[i].ID < todos[j].ID
	})

	return todos
}

// GetCompletedCount returns the number of completed todos
func (tm *TodoManager) GetCompletedCount() int {
	count := 0

	for _, todo := range tm.todos {
		if todo.Completed {
			count++
		}
	}

	return count
}

// GetPendingTodos returns all incomplete todos
func (tm *TodoManager) GetPendingTodos() []*Todo {
	var pending []*Todo

	for _, todo := range tm.todos {
		if !todo.Completed {
			pending = append(pending, todo)
		}
	}

	// Sort by ID
	sort.Slice(pending, func(i, j int) bool {
		return pending[i].ID < pending[j].ID
	})

	return pending
}

// GetCompletedTodos returns all completed todos
func (tm *TodoManager) GetCompletedTodos() []*Todo {
	var completed []*Todo

	for _, todo := range tm.todos {
		if todo.Completed {
			completed = append(completed, todo)
		}
	}

	// Sort by ID
	sort.Slice(completed, func(i, j int) bool {
		return completed[i].ID < completed[j].ID
	})

	return completed
}

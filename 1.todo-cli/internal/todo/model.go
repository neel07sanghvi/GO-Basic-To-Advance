package todo

import (
	"fmt"
	"time"
)

// Todo represents a single todo item
type Todo struct {
	ID          int        `json:"id"`
	Task        string     `json:"task"`
	Completed   bool       `json:"completed"`
	CreatedAt   time.Time  `json:"created_at"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
	// You need a pointer here because of how Go handles zero values and nullable fields:
	// time.Time has a zero value (January 1, year 1, 00:00:00 UTC)
	// If you used time.Time directly, you couldn't distinguish between "no completion time set" and "completed at the zero time"
	// *time.Time can be nil to represent "no completion time" or point to an actual time value

	// Without pointer - ambiguous
	// CompletedAt time.Time // Always has a value, even if meaningless

	// With pointer - clear semantics
	// CompletedAt *time.Time // nil = not completed, non-nil = completed at specific time
}

// String returns a formatted string representation of the todo
func (t *Todo) String() string {
	status := "[x]"

	if t.Completed {
		status = "[✓]"
	}

	completedInfo := ""

	if t.Completed && t.CompletedAt != nil {
		completedInfo = fmt.Sprintf(" (completed: %s)", t.CompletedAt.Format("2006-01-02 15:04"))
	}

	return fmt.Sprintf("%d. %s %s (created: %s)%s", t.ID, status, t.Task, t.CreatedAt.Format("2006-01-02 15:04"), completedInfo)
}

// MarkCompleted marks the todo as completed
func (t *Todo) MarkCompleted() {
	t.Completed = true
	now := time.Now()
	t.CompletedAt = &now

	// here make sure, you cannot do it like this
	// t.CompletedAt = &(time.Now())

	// They do look similar, but there's a subtle but important difference in Go's syntax rules.

	// Why &now Works:
	// now := time.Now()
	// t.CompletedAt = &now
	// now is a variable stored in memory
	// Variables have addresses, so you can take their address with &
	// ✅ This works perfectly

	// Why &(time.Now()) Doesn't Work:
	// t.CompletedAt = &(time.Now())  // ❌ Compile error
	// time.Now() is a function call expression that returns a value
	// Function return values are temporary - they don't have a guaranteed memory address
	// You cannot take the address of a function call result directly

	// The Compile Error:
	// cannot take the address of time.Now()
}

func (t *Todo) MarkIncomplete() {
	t.Completed = false
	t.CompletedAt = nil
}

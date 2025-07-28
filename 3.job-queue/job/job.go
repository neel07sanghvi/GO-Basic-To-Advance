package job

import (
	"fmt"
	"time"
)

type Job struct {
	ID      int
	Name    string
	Payload string
}

func (j *Job) Process() {
	fmt.Printf("Processing job %d: %s\n", j.ID, j.Name)

	time.Sleep(time.Millisecond * 500)

	fmt.Printf("Job %d completed: %s\n", j.ID, j.Payload)
}

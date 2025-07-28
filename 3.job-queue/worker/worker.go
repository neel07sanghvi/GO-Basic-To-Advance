package worker

import (
	"fmt"

	"githum.com/neel07sanghvi/job-queue/job"
)

type Worker struct {
	ID       int
	JobChan  chan job.Job
	QuitChan chan bool
}

func NewWorker(id int, jobChan chan job.Job) *Worker {
	return &Worker{
		ID:       id,
		JobChan:  jobChan,
		QuitChan: make(chan bool),
	}
}

func (w *Worker) Start() {
	fmt.Printf("Worder %d started\n", w.ID)

	for {
		select {
		case job := <-w.JobChan:
			fmt.Printf("Worker %d received job %d\n", w.ID, job.ID)
			job.Process()
			fmt.Printf("Worker %d finished job %d\n", w.ID, job.ID)

		case <-w.QuitChan:
			fmt.Printf("Worker %d stopping\n", w.ID)
			return
		}
	}
}

func (w *Worker) Stop() {
	w.QuitChan <- true
}

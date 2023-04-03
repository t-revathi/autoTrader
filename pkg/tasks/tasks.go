/*
Tasks
*/
package tasks

import (
	"time"
)

// base struct for tasks
type Task struct {
	Name      string
	Frequency time.Duration
	Handler   func() bool
	Laststart time.Time
}

//Factory method for task struct

func NewTask(
	name string,
	frequency time.Duration,
	handler func() bool,
	laststart time.Time,
) Task {
	return Task{Name: name,
		Frequency: frequency,
		Handler:   handler,
		Laststart: laststart,
	}
}

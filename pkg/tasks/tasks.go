package tasks

import (
	"time"
)

type Task struct {
	Name      string
	Frequency time.Duration
	Handler   func() bool
	Laststart time.Time
}

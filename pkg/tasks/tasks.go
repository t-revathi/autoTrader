package tasks

import (
	"fmt"
	"time"
)

type Task struct {
	Name      string
	Frequency time.Duration
	Handler   func() bool
	Laststart time.Time
}
type Tasks struct {
	taskList []Task
}

func AddTask(newTask Task) Tasks {
	var entry Tasks
	entry.taskList = append(entry.taskList, newTask)
	fmt.Println("added new task")
	return entry
}

func RunTasks(tasks Tasks) {
	for _, v := range tasks.taskList {

		t := time.Since(v.Laststart)
		fmt.Printf("Time diff:%v", t)
	}

}

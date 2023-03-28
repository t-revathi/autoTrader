package tasks

import (
	"context"
	"fmt"
	"time"
)
type TaskScheduler struct {
	taskList []Task
}

// factory method for schedulertasks
func NewTaskScheduler() *TaskScheduler {
	return &TaskScheduler{}
}

func (t *TaskScheduler) AddTask(newTask Task) error {

	t.taskList = append(t.taskList, newTask)
	//fmt.Printf("added new task %v", t.taskList)
	return nil
}

func (t *TaskScheduler) RunTasks(ctx context.Context, duration time.Duration) {

	ticker := time.NewTicker(duration)

	go func() {
		for {
			select {
			case <-ctx.Done():
				ticker.Stop()
			case <-ticker.C:
				fmt.Println("Ticks received")
				//fmt.Printf("added new task %v \n", t.taskList)
				for _, v := range t.taskList {

					v.Handler()
				}
			}
		}

	}()

}

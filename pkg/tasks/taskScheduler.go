package tasks

import (
	"context"
	"fmt"
	"time"
)
type SchedulerTasks struct {
	taskList []Task
}

func (t *SchedulerTasks) AddTask(newTask Task) error {

	t.taskList = append(t.taskList, newTask)
	//fmt.Printf("added new task %v", t.taskList)
	return nil
}

func (t *SchedulerTasks) RunTasks(ctx context.Context, duration time.Duration) {

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

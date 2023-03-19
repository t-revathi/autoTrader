package main

import (
	"autoTrader/pkg/tasks"
	service "autoTrader/service"
	"fmt"
	"time"
)

func main() {
	fmt.Printf("Hello")
	t := tasks.Task{
		Name:      "checkPrice",
		Frequency: 3000,
		Laststart: time.Now(),
		Handler:   service.GetPrices,
	}
	taskList := tasks.AddTask(t)
	tasks.RunTasks(taskList)
	fmt.Printf(t.Name)

}

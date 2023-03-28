package main

import (
	"autoTrader/pkg/tasks"
	service "autoTrader/service"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	taskObj := tasks.SchedulerTasks{}

	t := tasks.Task{
		Name:      "check Price",
		Frequency: 3,
		Laststart: time.Now(),
		Handler:   service.GetPrices,
	}

	if err := taskObj.AddTask(t); err != nil {
		fmt.Println("Could not add the task")
	}
	t1 := tasks.Task{
		Name:      "Buy Shares",
		Frequency: 3,
		Laststart: time.Now(),
		Handler:   service.BuyShares,
	}
	if err := taskObj.AddTask(t1); err != nil {
		fmt.Println("Could not add the task")
	}
	taskObj.RunTasks(ctx, time.Duration(3)*time.Second)
	//fmt.Printf(t.Name)

	go func() {
		for range interrupt {
			log.Println("\n❌ Interrupt received closing...")
			cancel()
		}
	}()

	<-ctx.Done()

}

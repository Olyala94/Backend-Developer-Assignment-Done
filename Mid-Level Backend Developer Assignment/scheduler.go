// scheduler.go

package main

import (
	"fmt"
	"time"
)

// ScheduleTask, bir iş birimini belirli bir zamanda veya aralıkta çalıştırmak için bir mekanizma ekler
func ScheduleTask(task Task, executionTime time.Time, taskQueue *TaskQueue) {
	go func() {
		time.Sleep(executionTime.Sub(time.Now()))
		taskQueue.Enqueue(task)
		fmt.Printf("İş %d zamanlanmış bir şekilde çalıştırılacak\n", task.ID)
	}()
}

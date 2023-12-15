// main.go

package main

import (
	"fmt"
	"sync"
)

func main() {
	taskQueue := NewTaskQueue(100)

	// Start multiple worker nodes
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			Worker(id, taskQueue)
		}(i)
	}

	// Enqueue tasks
	for i := 1; i <= 10; i++ {
		taskQueue.Enqueue(Task{ID: i, Payload: fmt.Sprintf("Task %d", i)})
	}

	wg.Wait()
}

// taskqueue.go

package main

import (
	"fmt"
	"time"
)

// Task represents a unit of work
type Task struct {
	ID       int
	Payload  string
	Attempts int
}

// TaskQueue represents the main task queue
type TaskQueue struct {
	queue chan Task
}

// NewTaskQueue creates a new task queue with a specified capacity
func NewTaskQueue(capacity int) *TaskQueue {
	return &TaskQueue{
		queue: make(chan Task, capacity),
	}
}

// Enqueue adds a task to the task queue
func (q *TaskQueue) Enqueue(task Task) {
	q.queue <- task
}

// Dequeue removes and returns a task from the task queue
func (q *TaskQueue) Dequeue() Task {
	return <-q.queue
}

// Worker represents a worker node that processes tasks
func Worker(id int, taskQueue *TaskQueue) {
	for {
		task := taskQueue.Dequeue()
		fmt.Printf("Worker %d processing task: %+v\n", id, task)
		// Implement task processing logic
		time.Sleep(time.Second) // Simulating task processing time
	}
}

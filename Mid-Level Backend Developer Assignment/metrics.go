// metrics.go

package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	taskExecutionTime = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "task_execution_time_seconds",
			Help:    "Histogram of task execution times",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"task_id"},
	)
)

func init() {
	prometheus.MustRegister(taskExecutionTime)
}

// InstrumentedWorker records task execution time
func InstrumentedWorker(id int, taskQueue *TaskQueue) {
	for {
		task := taskQueue.Dequeue()
		startTime := time.Now()

		fmt.Printf("Worker %d processing task: %+v\n", id, task)
		// Implement task processing logic
		time.Sleep(time.Second) // Simulating task processing time

		duration := time.Since(startTime).Seconds()
		taskExecutionTime.WithLabelValues(fmt.Sprintf("%d", task.ID)).Observe(duration)
	}
}

// StartMetricsServer starts a Prometheus metrics server
func StartMetricsServer() {
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":9090", nil)
}

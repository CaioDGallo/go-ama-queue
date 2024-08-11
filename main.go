package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

// Job represents a job to be processed
type Job struct {
	ID      int
	Payload string
}

// Worker represents a worker that processes jobs
type Worker struct {
	ID         int
	JobChannel chan Job
	WaitGroup  *sync.WaitGroup
}

// NewWorker creates a new worker
func NewWorker(id int, jobChannel chan Job, wg *sync.WaitGroup) Worker {
	return Worker{
		ID:         id,
		JobChannel: jobChannel,
		WaitGroup:  wg,
	}
}

// Start starts the worker to process jobs
func (w Worker) Start() {
	go func() {
		for job := range w.JobChannel {
			startTime := time.Now()
			fmt.Printf("Worker %d processing job %d with payload: %s\n", w.ID, job.ID, job.Payload)
			randInt := rand.Intn(10) + 1
			time.Sleep(time.Duration(randInt) * time.Second) // Simulate job processing time
			duration := time.Since(startTime)
			fmt.Printf("Worker %d finished job %d in %v\n", w.ID, job.ID, duration)
			w.WaitGroup.Done()
		}
	}()
}

func main() {
	const numWorkers = 5
	const numJobs = 5

	jobChannel := make(chan Job, numJobs)
	var wg sync.WaitGroup

	// Create and start workers
	for i := 1; i <= numWorkers; i++ {
		worker := NewWorker(i, jobChannel, &wg)
		worker.Start()
	}

	// Add jobs to the queue
	for j := 1; j <= numJobs; j++ {
		wg.Add(1)
		job := Job{ID: j, Payload: fmt.Sprintf("Payload %d", j)}
		jobChannel <- job
	}

	// Wait for all jobs to be processed
	wg.Wait()
	close(jobChannel)

	log.Println("All jobs processed")
}

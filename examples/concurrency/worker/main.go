// This sample program demonstrates how to use a buffered
// channel to work on multiple tasks with a predefined number
// of goroutines.
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Task struct {
	Id        int
	JobId     int
	Status    string
	CreatedOn time.Time
}

func (t *Task) Run() {

	sleep := rand.Int63n(1000)
	// Delaying the execution for the sake of example
	time.Sleep(time.Duration(sleep) * time.Millisecond)
	t.Status = "Completed"
}

// wg is used to wait for the program to finish.
var wg sync.WaitGroup

const noOfWorkers = 3

// main is the entry point for all Go programs.
func main() {
	// Create a buffered channel to manage the task queue.
	taskQueue := make(chan *Task, 10)

	// Launch goroutines to handle the work.
	// The worker1 process is distributing with the value of noOfWorkers.
	wg.Add(noOfWorkers)
	for gr := 1; gr <= noOfWorkers; gr++ {
		go worker(taskQueue, gr)
	}

	// Add Tasks into Buffered channel.
	for i := 1; i <= 10; i++ {
		taskQueue <- &Task{
			Id:        i,
			JobId:     100 + i,
			CreatedOn: time.Now(),
		}
	}

	// Close the channel
	close(taskQueue)

	// Wait for all the work to get done.
	wg.Wait()
}

// worker is launched as a goroutine to process Tasks from
// the buffered channel.
func worker(taskQueue <-chan *Task, workerId int) {
	// Schedule the call to Done method of WaitGroup.
	defer wg.Done()
	for v := range taskQueue {
		fmt.Printf("Worker%d: received request for Task:%d - Job:%d\n", workerId, v.Id, v.JobId)
		v.Run()
		// Display we finished the work.
		fmt.Printf("Worker%d: Status:%s for Task:%d - Job:%d\n", workerId, v.Status, v.Id, v.JobId)
	}
}

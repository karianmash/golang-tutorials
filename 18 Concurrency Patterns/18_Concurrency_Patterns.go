package main

import (
	"fmt"
	"sync"
	"time"
)

// Job represents the work to be done.
type Job struct {
	ID     int
	Number int
}

// Result represents the outcome of a job.
type Result struct {
	Job       Job
	Factorial int
}

// worker calculates the factorial of a number and sends the result.
func worker(id int, jobs <-chan Job, results chan<- Result) {
	for job := range jobs {
		fmt.Printf("Worker %d starting job %d\n", id, job.ID)
		f := 1
		for i := 1; i <= job.Number; i++ {
			f *= i
		}
		time.Sleep(time.Millisecond * 100) // Simulate work
		results <- Result{Job: job, Factorial: f}
		fmt.Printf("Worker %d finished job %d\n", id, job.ID)
	}
}

func main() {
	const numJobs = 10
	const numWorkers = 3

	jobs := make(chan Job, numJobs)
	results := make(chan Result, numJobs)

	var wg sync.WaitGroup

	// Start workers
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			worker(workerID, jobs, results)
		}(w)
	}

	// Send jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- Job{ID: j, Number: j}
	}
	close(jobs)

	// Wait for all workers to finish
	wg.Wait()
	close(results)

	// Collect results
	fmt.Println("\n--- Results ---")
	for r := range results {
		fmt.Printf("Job %d (Number %d): Factorial = %d\n", r.Job.ID, r.Job.Number, r.Factorial)
	}

	fmt.Println("All jobs processed.")
}

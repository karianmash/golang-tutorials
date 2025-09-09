# Go Concurrency Patterns: Worker Pool

This tutorial explores a common concurrency pattern in Go: the **Worker Pool**. A worker pool is a technique for distributing a set of tasks among a fixed number of worker goroutines. This pattern is useful for managing concurrent operations, limiting resource consumption, and processing a large number of jobs efficiently.

## What is a Worker Pool?

A worker pool consists of:
- A channel for sending jobs to workers.
- A channel for collecting results from workers.
- A fixed number of worker goroutines that read from the jobs channel, perform work, and write results to the results channel.
- A `sync.WaitGroup` to wait for all worker goroutines to complete their tasks.

## Example: Factorial Worker Pool

The `18_Concurrency_Patterns.go` file demonstrates a worker pool that calculates the factorial of numbers. We define `Job` and `Result` structs, a `worker` function, and orchestrate the job distribution and result collection in `main`.

```go
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
	Job      Job
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
```

## Running the Example

To run this example, navigate to the `18 Concurrency Patterns` directory in your terminal and execute:

```bash
go run 18_Concurrency_Patterns.go
```

This will demonstrate how jobs are distributed among the workers and how results are collected. The output will vary slightly due to the concurrent nature of goroutines, but you will see workers picking up jobs and reporting their completion, followed by the final results.

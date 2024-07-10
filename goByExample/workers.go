package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func startManager(workerCount, numJobs int, jobs chan<- int, results chan int, comms chan JobDone, doneCEO chan bool) {
	jobsDone := 0
	defer func() {
		close(jobs)
		close(results)
	}()
	for {
		comm := <-comms
		if !comm.done {
			fmt.Printf("Manager: Worker %v cut his finger on job %v\n", comm.workerId, comm.job)
			workerCount--
			jobs <- comm.job
			if workerCount == 0 {
				doneCEO <- false
				return
			}
		}
		if comm.done {
			jobsDone++
			if jobsDone == numJobs {
				fmt.Printf("How many workers are able to work: %v\n", workerCount)
				doneCEO <- true
				return
			}
		}
	}
}

type JobDone struct {
	workerId, job int
	done          bool
}

func startWorker(id int, jobs <-chan int, res chan<- int, comms chan JobDone, injuryChance float64) {
	for j := range /* don't put <- or you will receive always same jobs */ jobs {
		fmt.Printf("started worker %v on job %v\n", id, j)
		time.Sleep(time.Second)
		if rand.Float64() <= injuryChance {
			comms <- JobDone{id, j, false}
			return
		}
		res <- j * 2
		fmt.Printf("finished worker %v on job %v\n", id, j)
		comms <- JobDone{id, j, true}
	}
}

func main() {
	const numJobs = 30
	const numWorkers = 3
	const injuryChance = 0.1
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	comms := make(chan JobDone)
	doneCEO := make(chan bool, 1)

	for i := 1; i <= numJobs; i++ {
		jobs <- i
	}

	go startManager(numWorkers, numJobs, jobs, results, comms, doneCEO)
	for i := 1; i <= numWorkers; i++ {
		go startWorker(i, jobs, results, comms, injuryChance)
	}

	// block main until there is message to CEO
	msg := <-doneCEO
	if msg {
		fmt.Println("Job done")
	} else {
		fmt.Println("Job wasn't done")
	}

	res := make([]int, 0, numJobs)
	for j := range results {
		res = append(res, j)
	}
	fmt.Printf("Jobs: %v, Complete: %v, results: %v\n", numJobs, len(res), res)

	var stats runtime.MemStats
	runtime.ReadMemStats(&stats)
	fmt.Printf("Number of Goroutines: %d\n", runtime.NumGoroutine())
}

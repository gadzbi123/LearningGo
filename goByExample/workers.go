package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func startManager(workerCount, JOBS int, jobs chan<- int, results chan int, comms chan JobsDone, doneCEO chan bool) {
	jobsDone := 0
	for {
		select {
		case comm := <-comms:
			if !comm.done {
				fmt.Printf("Manager: Worker %v cut his finger on job %v\n", comm.workerId, comm.job)
				workerCount--
				if workerCount == 0 {
					doneCEO <- false
					close(jobs)
					close(results)
					return
				}
				go func() { jobs <- comm.job }()
			}
			if comm.done {
				jobsDone++
				if jobsDone == JOBS {
					fmt.Printf("How many workers are able to work: %v\n", workerCount)
					doneCEO <- true
					close(jobs)
					close(results)
					return
				}
			}
		default:
		}
	}
}

type JobsDone struct {
	workerId, job int
	done          bool
}

func startWorker(id int, jobs <-chan int, res chan<- int, comms chan JobsDone) {
	for j := range /* don't put <- or you will receive always same jobs */ jobs {
		fmt.Printf("started worker %v on job %v\n", id, j)
		time.Sleep(time.Second)
		// Chance on injury
		if rand.Int()%30 == 0 {
			// fmt.Printf("Worker %v cut his finger on job %v\n", id, j)
			comms <- JobsDone{id, j, false}
			return
		}
		res <- j * 2
		fmt.Printf("finished worker %v on job %v\n", id, j)
		comms <- JobsDone{id, j, true}
	}

}
func main() {
	const JOBS = 30
	const WORKERS = 3
	jobs := make(chan int, JOBS)
	results := make(chan int, JOBS)
	comms := make(chan JobsDone)
	doneCEO := make(chan bool, 1)

	go func() {
		for i := 1; i <= JOBS; i++ {
			jobs <- i
		}
	}()

	go startManager(WORKERS, JOBS, jobs, results, comms, doneCEO)
	for i := 1; i <= WORKERS; i++ {
		go startWorker(i, jobs, results, comms)
	}

f:
	for {
		select {
		case x := <-doneCEO:
			if x {
				fmt.Println("Job done")
				break f
			} else {
				fmt.Println("Job wasn't done")
				break f
			}
		default:
		}
	}

	res := make([]int, 0, JOBS)
	for j := range results {
		res = append(res, j)
	}
	fmt.Printf("RESULTS: %v\n", res)

	var stats runtime.MemStats
	runtime.ReadMemStats(&stats)
	fmt.Printf("Number of Goroutines: %d\n", runtime.NumGoroutine())
}

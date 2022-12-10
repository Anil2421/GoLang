package main

import (
	"fmt"
	"sync"
)

var wg = new(sync.WaitGroup)

func main() {
	// Create 2 channels jobs and results
	jobs := make(chan int, 10)
	results := make(chan int, 10)
	wg.Add(2)
	go PrintStuffs(jobs, results)
	for i := 0; i < 10; i++ {
		jobs <- i

	}
	close(jobs)
	wg.Wait()
	fmt.Println("THis is the sample program to show the 2 go channels")
}

func PrintStuffs(jobs chan int, results chan int) {
	//results <- jobs
	fmt.Println(<-jobs)
	fmt.Println(<-results)
	wg.Done()
}

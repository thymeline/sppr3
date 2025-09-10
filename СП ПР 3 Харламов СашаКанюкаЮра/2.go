package main

import (
	"fmt"
	"sync"
)

func main() {
	jobs := make(chan int, 10)
	results := make(chan int, 10)
	var wait sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wait.Add(1)
		go func() {
			defer wait.Done()
			for job := range jobs {
				results <- job * job
			}
		}()
	}

	for i := 1; i <= 10; i++ {
		jobs <- i
	}
	close(jobs)
	wait.Wait()
	close(results)

	for r := range results {
		fmt.Println(r)
	}
}
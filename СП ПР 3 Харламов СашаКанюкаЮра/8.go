package main

import (
    "fmt"
    "sync"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
    defer wg.Done()
    for job := range jobs {
        fmt.Printf("воркер %d заворкерил %d\n", id, job)
    }
}

func main() {
    const num = 3
    jobs := make(chan int)
    var wg sync.WaitGroup

    for i := 1; i <= num; i++ {
        wg.Add(1)
        go worker(i, jobs, &wg)
    }

    for i := 1; i <= 9; i++ {
        jobs <- i
    }
    
    close(jobs)
    wg.Wait()
}
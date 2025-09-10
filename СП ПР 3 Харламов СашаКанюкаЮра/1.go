package main

import (
	"fmt"
	"time"
	"sync"
)

func main() {
	var wait sync.WaitGroup
	wait.Add(1)

	go func() {
		defer wait.Done()
		for i := 1; i<= 5; i++ {
			fmt.Println(i)
			time.Sleep(1 * time.Second)
		}
	}()

	wait.Wait()
}
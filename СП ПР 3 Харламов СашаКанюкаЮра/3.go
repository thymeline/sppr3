package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(time.Millisecond * 200)
	requests := make(chan int, 15)

	for i:= 1; i <= 15; i++ {
		<- tick
		fmt.Println("обработка", i)
		requests <- i
	}
	close(requests)
}
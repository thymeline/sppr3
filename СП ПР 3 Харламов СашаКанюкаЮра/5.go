package main

import (
	"fmt"
	"os"
	"sync"
	"crypto/md5"
)

func main() {
	files := []string{"file1.txt", "file2.txt", "file3.txt"}
	var wait sync.WaitGroup
	
	for i := 0; i < len(files); i++ {
		wait.Add(1)
		go func(i int) {
			defer wait.Done()
			data, _ := os.ReadFile(files[i])
			fmt.Println(files[i], md5.Sum(data))
		}(i)
	}
	wait.Wait()
}
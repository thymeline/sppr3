package main

import (
	"fmt"
	"sync"
	"net/http"
)

func main() {
	urls := []string {"https://google.com", "https://yandex.ru", "https://github.com"}
	var wait sync.WaitGroup

	for i := 0; i < 3; i++ {
		wait.Add(1)
		go func() {
			defer wait.Done()
			resp, _ := http.Get(urls[i])
			fmt.Println(urls[i], resp.StatusCode)
			resp.Body.Close()
		}()
	}
	wait.Wait()
}
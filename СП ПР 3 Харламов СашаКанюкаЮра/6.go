package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func searchSource(ctx context.Context, name string) (string, error) {
	delay := time.Duration(rand.Intn(1000)) * time.Millisecond
	select {
	case <-time.After(delay):
		return fmt.Sprintln("данные", name), nil
	case <-ctx.Done():
		return "", ctx.Err()
	}
}

func firstResult(ctx context.Context, sources []string) (string, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	
	result := make(chan string, len(sources))
	
	for _, source := range sources {
		go func(s string) {
			if res, err := searchSource(ctx, s); err == nil {
				select {
				case result <- res:
					cancel()
				case <-ctx.Done():
				}
			}
		}(source)
	}

	select {
	case res := <-result:
		return res, nil
	case <-ctx.Done():
		return "", ctx.Err()
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	
	sources := []string{"бд1", "апи", "бд2"}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	
	if result, err := firstResult(ctx, sources); err == nil {
		fmt.Println(result)
	}
}
package _5_context_with_deadline

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1

		for {
			// Check if there is a cancel signal
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
				time.Sleep(1 * time.Second) // Simulate slow process
			}
		}
	}()

	return destination
}

func TestContextWithDeadline(t *testing.T) {
	fmt.Println("Total Goroutine:", runtime.NumGoroutine())

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	defer cancel()

	destination := CreateCounter(ctx)
	for n := range destination {
		fmt.Println("Counter:", n)
	}

	fmt.Println("Total Goroutine:", runtime.NumGoroutine())
	/*
		Total Goroutine: 2
		Counter: 1
		Counter: 2
		Counter: 3
		Counter: 4
		Counter: 5
		Total Goroutine: 2
		--- PASS: TestContextWithDeadline (5.00s)
	*/
}

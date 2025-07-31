package _4_context_with_timeout

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

func TestContextWithTimeout(t *testing.T) {
	fmt.Println("Total Goroutine:", runtime.NumGoroutine())

	// Make Context with Timeout
	parent := context.Background()

	// Cancel needs to be done when a process occurs faster (does not exceed the timeout)
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel() // If the process does not exceed the timeout, then we cancel the context manually (to make sure there are no processes / goroutines in the Background)

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
	*/
}

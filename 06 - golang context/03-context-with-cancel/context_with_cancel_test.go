package _3_context_with_cancel

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)

	// [Cause Goroutine Leak]
	//go func() {
	//	defer close(destination)
	//	counter := 1
	//	for {
	//		destination <- counter
	//		counter++
	//	}
	//}()

	// Implement Context with Cancel
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
			}
		}
	}()

	return destination
}

func TestContextWithCancel(t *testing.T) {

	fmt.Println("Total Goroutine:", runtime.NumGoroutine()) // Jumlah Goroutine

	// Make Context with Cancel
	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	destination := CreateCounter(ctx)

	for n := range destination {
		fmt.Println("Counter:", n)
		if n == 10 {
			break
		}
	}
	cancel() // Send Cancel Signal

	time.Sleep(2 * time.Second)

	fmt.Println("Total Goroutine:", runtime.NumGoroutine())
	// BEFORE USING CONTEXT
	/*
		Total Goroutine: 2
		Counter: 1
		Counter: 2
		Counter: 3
		Counter: 4
		Counter: 5
		Counter: 6
		Counter: 7
		Counter: 8
		Counter: 9
		Counter: 10
		Total Goroutine: 3 // harusnya balik lagi ke 2 (awal) [Goroutine Leak]
	*/

	// AFTER USING CONTEXT
	/*
		Total Goroutine: 2
		Counter: 1
		Counter: 2
		Counter: 3
		Counter: 4
		Counter: 5
		Counter: 6
		Counter: 7
		Counter: 8
		Counter: 9
		Counter: 10
		Total Goroutine: 2
	*/
}

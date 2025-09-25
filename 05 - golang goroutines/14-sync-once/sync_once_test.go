package _14_sync_once

import (
	"fmt"
	"sync"
	"testing"
)

/* ===== INTRO =====
- Once adalah fitur di Golang yang bisa kita gunakan untuk MEMASTIKAN bahwa sebuah function di-eksekusi HANYA SEKALI.
- Jadi berapa banyakpun goroutine YANG MENGAKSES, bisa DIPASTIKAN bahwa goroutine yang PERTAMA yang bisa mengeksekusi function-nya.
- Goroutine yang lain akan DIHIRAUKAN, artinya function tidak akan dieksekusi lagi.
*/

var counter = 0

func OnlyOnce() {
	counter++
}

func TestOnce(t *testing.T) {
	var once sync.Once
	var group sync.WaitGroup

	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			once.Do(OnlyOnce)
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Counter:", counter)

	/*
		=== RUN   TestOnce
		Counter: 1
		--- PASS: TestOnce (0.00s)
		PASS
	*/
}

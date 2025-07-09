package _10_sync_mutex

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/* ===== INTRO =====
- Mutex (Mutual Exclusion)
- Untuk mengatasi masalah Race Condition, di Golang terdapat sebuah struct bernama sync.Mutex
- Mutex bisa digunakan untuk melakukan locking dan unlocking, dimana ketika kita melakukan LOCKING terhadap
  mutex, maka TIDAK ADA yang bisa melakukan locking lagi sampai kita melakukan UNLOCK
- Dengan demikian, jika ada BEBERAPA GOROUTINE melakukan LOCK terhadap Mutex, maka HANYA 1 goroutine yang
  DIPERBOLEHKAN, setelah goroutine tersebut melakukan UNLOCK, baru goroutine selanjutnya diperbolehkan
  melakukan LOCK lagi
- Ini sangat cocok sebagai solusi ketika ada masalah race condition yang sebelumnya kita hadapi
*/

func TestSyncMutex(t *testing.T) {
	x := 0
	var mutex sync.Mutex

	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("Counter:", x)
	/*
		=== RUN   TestSyncMutex
		Counter: 100000
		--- PASS: TestSyncMutex (5.00s)
		PASS
	*/
}

package _17_sync_cond

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/* ===== INTRO =====
- Cond adalah implementasi LOCKING berbasis KONDISI.
- Cond membutuhkan LOCKER (bisa menggunakan Mutext / RWMutex) utk implementasi locking-nya, namun berbeda
  dengan Locker biasanya, di Cond terdapat function Wait() untuk menunggu apakah PERLU MENUNGGU atau TIDAK.
- Function Signal() bisa digunakan untuk memberi tahu sebuah goroutine agar TIDAK PERLU MENUNGGU LAGI, sedangnkan
  function Broadcast() digunakan unutk memberi tahu SEMUA GOROUTINE AGAR TIDAK PERLU MENUNGGU LAGI.
- Untuk mmebuat Cond, kita bisa menggunakan function sync.NewCond(Locker).

*/

var (
	cond  = sync.NewCond(&sync.Mutex{})
	group = sync.WaitGroup{}
)

func WaitCond(value int) {
	cond.L.Lock()
	cond.Wait()
	fmt.Println("DONE:", value)
	cond.L.Unlock()
	group.Done()
}

func TestCond(t *testing.T) {
	for i := 0; i < 10; i++ {
		group.Add(1)
		go WaitCond(i)
	}

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			cond.Signal()
		}
	}()

	group.Wait()

	/*
			=== RUN   TestCond
		DONE: 1
		DONE: 3
		DONE: 2
		DONE: 4
		DONE: 5
		DONE: 6
		DONE: 7
		DONE: 8
		DONE: 9
		DONE: 0
		--- PASS: TestCond (10.00s)
		PASS
	*/
}

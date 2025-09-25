package _18_atomic

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

/* ===== INTRO =====
- Golang memiliki package yang bernama sync/atomic.
- Atomic merupakan package yang digunakan untuk menggunakan DATA PRIMITIVE secara AMAN pada proses CONCURRENT.
- Contohnya sebelumnya kita telah menggunakan Mutex untuk melakukan Locking ketika ingin menaikkan angka di coutner.
  Hal ini sebenarnya bisa digunakan menggunakan Atomic Package.
- Ada banyak sekali function di atomic package, kita bisa eksplore sendiri di halaman dokumentasinya.
- https://golang.org/pkg/sync/atomic/
*/

func TestAtomic(t *testing.T) {
	var group sync.WaitGroup
	var counter int64 = 0

	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			for j := 0; j < 100; j++ {
				atomic.AddInt64(&counter, 1)
			}
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Counter:", counter)
	/*
		=== RUN   TestAtomic
		Counter: 10000
		--- PASS: TestAtomic (0.00s)
		PASS
	*/
}

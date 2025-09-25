package _13_sync_waitgroup

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/* ===== INTRO =====
- WaitGroup adalah fitur yang bisa digunakan untuk menunggu sebuah proses selesai dilakukan
- Hal ini kadang diperlukan, misal kita ingin menjalankan beberapa proses menggunakan goroutine,
  tapi kita INGIN SEMUA PROSES SELESAI terlebih dahulu SEBELUM APLIKASI SELESAI
- Kasus seperti ini bisa menggunakan WaitGroup
- Untuk menandai bahwa ada proses goroutine

*/

func RunAsynchronous(group *sync.WaitGroup) {
	defer group.Done()
	group.Add(1)

	fmt.Println("Hello")
	time.Sleep(1 * time.Second)
}

func TestSyncWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsynchronous(group)
	}
	group.Wait()

	fmt.Println("Complete!!")

	/*
	   === RUN   TestSyncWaitGroup
	   Hello
	   ...
	   Hello
	   Hello
	   Complete!!
	   --- PASS: TestSyncWaitGroup (1.00s)
	   PASS
	*/
}

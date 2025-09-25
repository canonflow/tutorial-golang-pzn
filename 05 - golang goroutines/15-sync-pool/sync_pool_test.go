package _15_sync_pool

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"
)

/* ===== INTRO =====
- Pool adalah implementasi design pattern bernama OBJECT POOL PATTERN.
- Sederhananya, design pattern Pool ini digunakan untuk menyimpan data, selanjutnya untuk menggunakan datanya
  kita bisa mengambil dari Pool, dan setelah selesai menggunakan datanya, kita bisa menyimpan kembali ke Pool-nya.
- Implementasi Pool di Golang ini sudah aman dari problem race condition.
*/

func TestSyncPool(t *testing.T) {
	var pool sync.Pool
	pool.Put("Nathan")
	pool.Put("Garzya")
	pool.Put("Santoso")

	// Kalau mau membuat data pool secara otomatis
	// pool := sync.Pool{
	// 	New: func() interface{} {
	// 		return "New"
	// 	},
	// }

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println("Counter: " + strconv.Itoa(i) + " - Data: " + data.(string))
			pool.Put(data)
		}()
	}

	time.Sleep(time.Second * 3)

	/*
		=== RUN   TestSyncPool
		Counter: 0 - Data: Garzya
		Counter: 3 - Data: Garzya
		Counter: 2 - Data: Garzya
		Counter: 5 - Data: Garzya
		Counter: 4 - Data: Garzya
		Counter: 6 - Data: Garzya
		Counter: 7 - Data: Garzya
		Counter: 8 - Data: Garzya
		Counter: 9 - Data: Nathan
		Counter: 1 - Data: Santoso
		--- PASS: TestSyncPool (3.00s)
		PASS
	*/
}

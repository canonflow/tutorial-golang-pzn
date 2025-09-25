package _21_gomaxprocs

import (
	"fmt"
	"runtime"
	"testing"
)

/* ===== INTRO =====
- Sebelumnya di awal kita sudah bahas bahwa goroutine itu sebenarnya DIJALANKAN di dalam THREAD.
- Pertanyaannya, seberapa banyak Thread YANG ADA di Golang ketika aplikasi kita BERJALAN?
- Untuk mengetahui BERAPA JUMLAH Thread, kita bisa menggunakan GOMAXPROCS, yaitu sebuah function di package
  runtime yang bisa kita gunakan untuk MENGUBAH JUMLAH Thread atau MENGAMBIL JUMLAH Thread.
- Secara DEFAULT, jumlah Thread di Golang itu sebanyak jumlah CPU di komputer kita
- Kita juga bisa melihat berapa CPU kita dengan menggunakan function runtime.NumCpu().
*/

/* ===== PERINGATAN =====
- Menambah jumlah Thread TIDAK BERARTI membuat aplikasi kita menjadi lebih cepat.
- Karena pada saat yang SAMA, 1 CPU hanya AKAN MENJALANKAN 1 Goroutine dengan 1 Thread.
- Oleh karena ini, jika kita INGIN MENAMBAH THROUGHPUT aplikasi, disarankan lakukan VERTICAL SCALING
  (dengan menambah jumlah CPU) atau HORIZONTAL SCALING (menambah node baru).
*/

func TestThreadCount(t *testing.T) {
	totalCpu := runtime.NumCPU()
	fmt.Println("CPU Count:", totalCpu)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Thread Count:", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Goroutine Count:", totalGoroutine)

	/*
		=== RUN   TestThreadCount
		CPU Count: 16
		Thread Count: 16
		Goroutine Count: 2
		--- PASS: TestThreadCount (0.00s)
		PASS
	*/
}

func TestChangeThreadCount(t *testing.T) {
	totalCpu := runtime.NumCPU()
	fmt.Println("CPU Count:", totalCpu)

	runtime.GOMAXPROCS(20)
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Thread Count:", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Goroutine Count:", totalGoroutine)

	/*
		=== RUN   TestChangeThreadCount
		CPU Count: 16
		Thread Count: 20
		Goroutine Count: 2
		--- PASS: TestChangeThreadCount (0.00s)
		PASS
	*/
}

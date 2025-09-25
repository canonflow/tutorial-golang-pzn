package _20_ticker

import (
	"fmt"
	"testing"
	"time"
)

/* ===== INTRO =====
- Ticker adalah representasi kejadian YANG BERULANG.
- Ketika waktu ticker sudah EXPIRE, maka event akan DIKIRIM ke dalam Channel.
- Untuk membuat ticker, kita bisa menggunakan time.NewTicker(duration).
- Untuk menghentikan ticker, kita bisa menggunakan Ticker.Stop().
*/

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)

	for tick := range ticker.C {
		fmt.Println(tick)
	}

	/*
		=== RUN   TestTicker
		2025-09-25 20:51:53.0321819 +0700 +07 m=+1.001135501
		2025-09-25 20:51:54.0321819 +0700 +07 m=+2.001135501
		2025-09-25 20:51:55.0321819 +0700 +07 m=+3.001135501
		2025-09-25 20:51:56.0321819 +0700 +07 m=+4.001135501
		2025-09-25 20:51:57.0321819 +0700 +07 m=+5.001135501
		2025-09-25 20:51:58.0321819 +0700 +07 m=+6.001135501
		2025-09-25 20:51:59.0321819 +0700 +07 m=+7.001135501
	*/
}

/* ===== Tick() =====
- Kadang kita tidak butuh data Ticker-nya, kita hanya butuh channel-nya saja.
- Jika demikian, kita bisa menggunakan function timer.Tick(duration), function ini tidak akan
  mengembalikan Ticker, hanya mengembalikan channel timernya saja.
*/

func TestTick(t *testing.T) {
	channel := time.Tick(1 * time.Second)

	for tick := range channel {
		fmt.Println(tick)
	}

	/*
		=== RUN   TestTick
		2025-09-25 20:54:03.3178687 +0700 +07 m=+1.000840601
		2025-09-25 20:54:04.3178687 +0700 +07 m=+2.000840601
		2025-09-25 20:54:05.3178687 +0700 +07 m=+3.000840601
		2025-09-25 20:54:06.3178687 +0700 +07 m=+4.000840601
		2025-09-25 20:54:07.3178687 +0700 +07 m=+5.000840601
		2025-09-25 20:54:08.3178687 +0700 +07 m=+6.000840601
		2025-09-25 20:54:09.3178687 +0700 +07 m=+7.000840601
		2025-09-25 20:54:10.3178687 +0700 +07 m=+8.000840601
		2025-09-25 20:54:11.3178687 +0700 +07 m=+9.000840601
	*/
}

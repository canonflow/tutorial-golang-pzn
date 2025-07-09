package _6_range_channel

import (
	"fmt"
	"strconv"
	"testing"
)

/* ===== INTRO =====
- Kadang - kadang ada kasus sebuah channel dikirim data secara terus menerus oleh pengirim
- Dan kadang tidak jelas kapan channel tersebut akan berhenti menerima data
- Salah satu yang bisa kita lakukan adalah dengan menggunakan perulangan range ketika menerima data dari channel
- Ketika sebuah channel di close(), maka secara otomatis perulangan tersebut akan BERHENTI
- Ini lebih sederhana dari pada kita melakukan pengecekan channel secara manual
*/

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
		close(channel) // Harus di close, kalau nggak nanti deadlock karena for range akan terus mengecek channel untuk menerima data
	}()

	for data := range channel {
		fmt.Println(data)
	}

	fmt.Println("===== SELESAI =====")
	/*
		=== RUN   TestRangeChannel
		Perulangan ke 0
		Perulangan ke 1
		Perulangan ke 2
		Perulangan ke 3
		Perulangan ke 4
		Perulangan ke 5
		Perulangan ke 6
		Perulangan ke 7
		Perulangan ke 8
		Perulangan ke 9
		===== SELESAI =====
		--- PASS: TestRangeChannel (0.00s)
		PASS
	*/
}

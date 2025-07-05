package _5_buffered_channel

import (
	"fmt"
	"testing"
	"time"
)

/* ===== INTRO =====
- Secara default, channel hanya bisa menerima 1 data
- Artinya jika kita menambahkan data ke-2, maka kita akan diminta menunggu sampai data ke-1 ada yg mengambil
- Kadang2 ada kasus dimana PENGIRIM LEBIH CEPAT DIBANDING PENERIMA, dalam hal ini jika kita menggunakan channel
  maka otomatis pengirim akan ikut lambat juga.
- Untungnya ada BufferedChannel, yaitu buffer yg bisa digunakan untuk MENAMPUNG data antrian di channel

===== BUFFER CAPACITY =====
- Kita bebas memasukka berapa jumlah kapasitas antrian di dalam buffer
- Jika kita set misal 5, artinya kita bisa menerima 5 data di buffer
- Jika kita mengirim data ke-6, maka kita diminta untuk menunggu buffer sampai ada yg kosong
- Ini cocok sekali ketika memang goroutine yg menerima data LEBIH LAMBAT dari yang mengirim data
- Kalau data nggak ada yg ambil tidak apa2, soalnya dimasukkan ke buffer. Kecuali kalau melebih kapasitas
- Kalau mau mengambil, tinggal <-channel sebanyak data yang dikirimkan. Kalau tidak cocok, akan deadlock
*/

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	go func() {
		channel <- "Nathan"
		channel <- "Garzya"
		channel <- "Santoso"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)

	fmt.Println("SELESAI")
	/*
		=== RUN   TestBufferedChannel
		Nathan
		Garzya
		Santoso
		SELESAI
		--- PASS: TestBufferedChannel (2.00s)
	*/
}

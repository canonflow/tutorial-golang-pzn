package _2_channel

import (
	"fmt"
	"testing"
	"time"
)

/* ===== INTRO ======
- Channel adalah tempat komunikasi secara synchronous yg bisa dilakukan oleh goroutine
- Di channel terdapat pengirim dan penerima, biasanya pengirim dan penerima adalah goroutine yang BERBEDA
- Saat melakukan pengiriman ke Channel, goroutine akan TER-BLOC, sampai ada yang menerima data tersebut
- Maka dari itu, channel disebut sebagai alat komunikasi synchronous (blocking)
- Channel cocok sekali sebagai alternatif seperti mekanisme async-await yg terdapat di beberapa
  bahasa pemrograman lain

===== KARAKTERISTIK =====
- Secara default channel hanya bisa menampung 1 data, jika kita ingin menambahkan data lagi,
  harus menunggu data yang ada di channel diambil
- Channel hanya bisa menerima SATU JENIS DATA
- Channel bisa diambil dari lebih dari 1 goroutine
- Channel harus DI CLOSE jika tidak digunakan, atau bisa menyebabkan memory leak

===== HOW TO MAKE =====
- Channel di Golang direpresentasikan dengan tipe data CHAN
- Untuk membuat channel sangat mudah, kita bisa menggunakan make() mirip ketika membuat map
- Namun saat pembuatan channel, kita harus tentukan tipe data apa yang bisa dimasukkan ke dalam channel tsb

===== SEND and RECEIVE DATA =====
- Send: channel <- data
- Receive: data <- channel
- Close: close()
*/

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Hello"
		fmt.Println("Selesai mengirim data ke channel")
	}()

	data := <-channel
	fmt.Println(data)
	/*
		=== RUN   TestCreateChannel
		Hello
		Selesai mengirim data ke channel
		--- PASS: TestCreateChannel (2.00s)
		PASS
	*/
}

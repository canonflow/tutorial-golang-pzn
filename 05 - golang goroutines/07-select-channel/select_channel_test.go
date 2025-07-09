package _7_select_channel

import (
	"fmt"
	"testing"
	"time"
)

/* ===== INTRO =====
- Kadang ada kasus dimana kita membuat BEBERAPA CHANNEL dan menjalankan BEBERAPA GOROUTINE
- Lalu kita ingin mendapatkan data dari semua channel tersebut
- Untuk melakukan hal tersebut, kita bisa menggunakan select di Golang
- Dengan select channel, kita bisa memilih DATA TERCEPAT dari beberapa channel, jika data datang SECARA BERSAMAAN
  di beberapa channel, maka dipilih secara RANDOM
*/

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Nathan"
}

func TestSelectMultipleChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel 2", data)
			counter++
		}

		if counter == 2 {
			break
		}
	}
	/*
		=== RUN   TestSelectMultipleChannel
		Data dari channel 2 Nathan
		Data dari channel 1 Nathan
		--- PASS: TestSelectMultipleChannel (2.00s)
		PASS
	*/
}

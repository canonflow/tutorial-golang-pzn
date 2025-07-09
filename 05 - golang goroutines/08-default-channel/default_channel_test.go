package _8_default_channel

import (
	"fmt"
	"testing"
	"time"
)

/* ===== INTRO =====
- Apa yang terjadi jika kita melakukan select terhadap channel yang ternyata TIDAK ADA DATANYA?
- Maka kita akan menunggu sampai data ada
- Kadang mungkin kita ingin melakukan sesuatu jika misal SEMUA CHANNEL TIDAK ADA DATANYA ketika
  kita melakukan select channel
- Dalam select, kita bisa menambahkan DEFAULT, dimana ini akan dieksekusi jika memang di SEMUA CHANNEL
  yang kita selecta TIDAK ADA DATANYA
*/

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Nathan"
}

func TestDefaultChannel(t *testing.T) {
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
		default:
			fmt.Println("Menunggu Data ...")
		}

		if counter == 2 {
			break
		}
	}
	/*
		...
		Menunggu Data ...
		Menunggu Data ...
		Menunggu Data ...
		Menunggu Data ...
		Menunggu Data ...
		Data dari channel 2 Nathan
		Data dari channel 1 Nathan
		--- PASS: TestDefaultChannel (2.00s)
		PASS
	*/
}

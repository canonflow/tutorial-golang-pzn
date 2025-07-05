package _4_channel_in_out

import (
	"fmt"
	"testing"
	"time"
)

/* ===== INTRO =====
- Saat mengirim channel sbg parameter, isi functiont tsb bisa mengirim dan menerima data dari channel tsb
- Kadang kita ingin memberi tahu thd function, misal bahwa channel tsb hanya digunakan untuk mengirimkan data,
  atau hanya dapat digunakan utk menerima data
- Hal ini bisa kita lakukan di parameter dengan cara menandai apakah channel ini digunakan untuk
  in (mengirim data ke channel) atau out (menerima data dari channel)
*/

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Nathan"
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
	/*
		=== RUN   TestInOutChannel
		Nathan
		--- PASS: TestInOutChannel (5.00s)
		PASS
	*/
}

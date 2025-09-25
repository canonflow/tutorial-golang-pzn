package _19_timer

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/* ===== INTRO =====
- Timer adalah representasi satu kejadian.
- Ketika waktu timer sudah EXPIRE, maka EVENT akan DIKIRIM ke dalam channel.
- Untuk membuat Timer, kita bisa menggunakan time.NewTimer(duration).
*/

func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)

	fmt.Println(time.Now())
	time := <-timer.C
	fmt.Println(time)

	/*
		=== RUN   TestTimer
		2025-09-25 20:45:17.2988242 +0700 +07 m=+0.001543601
		2025-09-25 20:45:22.2988242 +0700 +07 m=+5.001543601
		--- PASS: TestTimer (5.00s)
		PASS
	*/
}

/* ===== TIME AFTER =====
- Kadang kita HANYA BUTUH Channel-nya saja, tidak membutuhkan data Timer-nya.
- Untuk melakukan hal itu, kita bisa menggunakan function time.After(duration)
*/

func TestAfter(t *testing.T) {
	channel := time.After(1 * time.Second)

	tick := <-channel
	fmt.Println(tick)

	/*
		=== RUN   TestAfter
		2025-09-25 20:46:58.2358806 +0700 +07 m=+1.001094201
		--- PASS: TestAfter (1.00s)
		PASS
	*/
}

/* ===== AFTER FUNC =====
- Kadang ada kebutuhan kita ingin menjalankan sebuah function DENGAN DELAY waktu tertentu.
- Kita bisa memanfaatkan Timer dengan menggunakan function time.AfterFunc().
- Kita TIDAK PERLU LAGI menggunakan channel-nya, cukup KIRIMKAN function yang AKAN DIPANGGIL ketika Timer mengirimkan kejadiannya.
*/

func TestAfterFunc(t *testing.T) {
	group := sync.WaitGroup{}
	group.Add(1)

	time.AfterFunc(2*time.Second, func() {
		fmt.Println("Execute after 2 seconds")
		group.Done()
	})

	group.Wait()

	/*
		=== RUN   TestAfterFunc
		Execute after 2 seconds
		--- PASS: TestAfterFunc (2.00s)
		PASS
	*/
}

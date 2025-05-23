package main

import "fmt"

func endApplication() {
	fmt.Println("End App")
	// CARA YANG BENAR
	message := recover()
	fmt.Println("Terjadi error", message)
}

func runApps(err bool) {
	defer endApplication()
	if err {
		panic("Upps, Error :)") // AKAN BERHENTI TOTAL
	}

	// CARA YG SALAH
	//message := recover()
	//fmt.Println("Terjadi panic", message)
}

func main() {
	/* ===== INTRO =====
	- Recover adalah sebuah function yg digunakan untuk menangkap data Panic
	- Dengan Recover proses Panic akan TERHENTI, sehingga program AKAN TETAP BERJALAN
	- Recover perlu dipanggil di dalam function yang akan didefer
	*/
	runApps(true)
	/*
		End App
		Terjadi error Upps, Error :)
	*/
	fmt.Println("Nathan Garzya") // Nathan Garzya
}

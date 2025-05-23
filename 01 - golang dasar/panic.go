package main

import "fmt"

func endApp() {
	fmt.Println("End App")
}

func runApp(err bool) {
	defer endApp()
	if err {
		panic("Upps, Error :)")
	}
}

func main() {
	/* ===== INTRO =====
	- Panic adalah sbh function yg bisa kita gunakan untuk MENGHENTIKAN PROGRAM
	- Function Panic biasanya dipanggil ketika terjadi suatu panic saat program kita berjalan
	- Saat dipanggil, function akan BERHENTI, tetapi Defer Function TETAP AKAN DIEKSEKUSI
	*/

	runApp(false) // End App
	runApp(true)
	/*
		End App
		panic: Upps, Error :)

		goroutine 1 [running]:
		main.runApp(0x0?)
		        D:/Code/Course/Programmer Zaman Now/tutorial-golang/01 - golang dasar/panic.go:12 +0x52
		main.main()
		        D:/Code/Course/Programmer Zaman Now/tutorial-golang/01 - golang dasar/panic.go:24 +0x1f
		exit status 2
	*/

}

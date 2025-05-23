package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication() {
	defer logging() // Akan tetap dijalankan di akhir eksekusi walaupun terdapat error
	fmt.Println("Run Application")
}

func main() {
	runApplication()
	/*
		Run Application
		Selesai Memanggil Function
	*/
}

package main

import "fmt"

func main() {
	fmt.Println("Nathan")
	fmt.Println("Nathan Garzya")
	fmt.Println("Nathan Garzya Santoso")

	fmt.Println(len("Nathan"))      // Mengambil panjang string
	fmt.Println("Nathan Garzya"[0]) // Akan mengambilkan huruf ke-i dalam bentuk byte / ASCII
	fmt.Println("Nathan Garzya Santoso"[1])
}

package main

import "fmt"

func main() {
	// NoKTP adalah tipe data baru, tetapi sebenarnya berupa alias dari tipe data string
	type NoKTP string

	var nathanKTP NoKTP = "11111111"
	var contoh string = "22222222"
	var contohKTP NoKTP = NoKTP(contoh) // Konversi string ke NoKTP

	fmt.Println(nathanKTP)
	fmt.Println(contohKTP)
}

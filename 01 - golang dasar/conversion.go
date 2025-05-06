package main

import "fmt"

func main() {
	// ===== NUMBER =====
	var nilai32 int32 = 32768

	var nilai64 int64 = int64(nilai32)

	// Perlu diperhatikan jika ingin konversi number ke bawahnya
	// Pastikan nilai sebelumnya dalam jangkauan tipe data yg baru
	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai32) // 32768
	fmt.Println(nilai64) // 32768
	fmt.Println(nilai16) // -32678 -> melebihi kapasitas (number overflow). Setelah mencapai max, akan kembali ke min lalu dihitung kembali

	// ===== STRING =====
	var name string = "Nathan"
	var n uint8 = name[0]
	var nString = string(n)

	fmt.Println(name)    // Nathan
	fmt.Println(n)       // 78
	fmt.Println(nString) // N
}

package main

/* ===== BLANK IDENTIFIER =====
- Kadang kita hanya ingin menjalankan init function di package tanpa harus mengeksekusi salah 1 function yg ada di package
- Secara default, golang akan komplen ketika ada package yg diimport namun tidak digunakan
- Untuk menangani hal tsb, kita bisa menggunakan BLANK IDENTIFIER (_) sblm nama package ketika melakukan import

*/

import (
	"fmt"
	"golang-dasar/database"
	_ "golang-dasar/internal"
)

func main() {
	fmt.Println(database.GetDatabase()) // MySQL
}

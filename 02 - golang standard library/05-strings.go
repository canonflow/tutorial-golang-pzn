package main

import (
	"fmt"
	"strings"
)

func main() {
	/* ===== INTRO =====
	- Package strings adalah package yg berisikan function - function utk memanipulasi tipe data String
	- Beberapa Function
		- Trim(string, cutset): memotong cutset di awal dan di akhir string
		- ToLower(string): Membuat semua karakter string menjadi lower case
		- ToUpper(string): Membuat semua karakter string menjadi upper case
		- Split(string, separator): Memotong string berdasarkan separator
		- Contains(string, search): Mengecek apakah string mengandung string lain
		- ReplaceAll(string, from, to): Mengubah semua string dari from ke to
	*/
	fmt.Println(strings.Contains("Nathan Garzya", "Nathan"))                               // true
	fmt.Println(strings.Split("Nathan Garzya Santoso", " "))                               // [Nathan Garzya Santoso]
	fmt.Println(strings.ToLower("Nathan Garzya Santoso"))                                  // nathan garzya santoso
	fmt.Println(strings.ToUpper("Nathan Garzya Santoso"))                                  // NATHAN GARZYA SANTOSO
	fmt.Println(strings.Trim("                  Nathan Garzya Santoso            ", " "))  // Nathan Garzya Santoso
	fmt.Println(strings.ReplaceAll("Nathan Garzya Nathan Santoso", "Nathan", "Canonflow")) // Canonflow Garzya Canonflow Santoso

}

package main

import (
	"fmt"
	"slices"
)

func main() {
	/* ===== INTRO =====
	- Di golang versi terbaru, terdapat fitur bernama Generic, fitur ini akan kita bahas khusus di kelas Golang Generic
	- Fitur Generic ini membuat kita bisa membuat parameter dengan tipe yg bisa berubah - ubah, tanpa harus menggunakan
	  interface kosong / any
	- Salah satu package yg menggunakan fitur Generic ini adalah package slices
	- Package slices ini digunakan untuk memanipulasi data di slice
	*/
	names := []string{
		"John", "Paul", "George", "Ringo",
	}
	values := []int{100, 95, 80, 90}

	fmt.Println(slices.Min(names))                // George
	fmt.Println(slices.Min(values))               // 80
	fmt.Println(slices.Max(names))                // Ringo
	fmt.Println(slices.Max(values))               // 100
	fmt.Println(slices.Contains(names, "Nathan")) // False
	fmt.Println(slices.Index(names, "Nathan"))    // -1
	fmt.Println(slices.Index(names, "Paul"))      // 1
}

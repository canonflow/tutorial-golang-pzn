package main

import "fmt"

type House struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(house House) {
	house.City = "Indonesia"
}

func ChangeCountryToIndonesiaPointer(house *House) {
	house.City = "Indonesia"
}

func main() {
	/* ===== INTRO =====
	- Saat kita membuat parameter di function, secara default adalah PASS BY VALUE
	- Jika kita ingin mengubah data di dalam function, data yg aslinya tidak akan pernah berubah.
	  Oleh karena itu, variable menjadi aman karena tidak akan bisa diubah
	- Namun jika kita ingin mengubah data asli pada parameter, kita bisa menggunakan pointer
	- Untuk menjadikan sebuah parameter sbg pointer, kita bisa menggunakan operator * di parameter
	*/
	address1 := House{}
	ChangeCountryToIndonesia(address1)
	fmt.Println(address1) // Tidak Berubah

	ChangeCountryToIndonesiaPointer(&address1)
	fmt.Println(address1) // {Indonesia  }
}

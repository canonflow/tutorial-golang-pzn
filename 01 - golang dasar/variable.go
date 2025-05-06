package main

import "fmt"

func main() {
	// ===== Hanya deklarasi saja, maka wajib menyebutkan tipe datanya =====
	var name string

	name = "Nathan"
	fmt.Println(name) // Nathan

	name = "Garzya"
	fmt.Println(name) // Garzya

	// ===== Jika inisialisasi langsung, maka tidak wajib menyebutkan tipe datanya =====
	var street = "Raya Tenggilis"
	fmt.Println(street) // Raya Tenggilis

	// ===== Jika inisialisasi langsung, tidak wajib juga menggunakan keyword var =====
	hobby := "Swimming"
	fmt.Println(hobby) // Swimming

	hobby = "Running"  // Jika ingin mengubah nilainya, tidak perlu memakai :
	fmt.Println(hobby) // Running

	// ===== Deklarasi multiple variable =====
	var (
		firstName = "Nathan Garzya"
		lastName  = "Santoso"
	)

	fmt.Println(firstName, lastName) // Nathan Garzya Santoso
}

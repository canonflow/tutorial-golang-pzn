package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	/* ===== INTRO =====
	- Secara default di Golang semua variable itu di passing by value, bukan by reference
	- Artinya, jika kita mengirim sebuah variable ke dalam function, method, atau variable lain
	  sebenarnya yg dikirim adlah duplikasi value-nya
	*/
	// PASS BY VALUE
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := address1 // Copy Value

	address2.City = "Bandung"
	fmt.Println(address1) // {Subang Jawa Barat Indonesia}
	fmt.Println(address2) // {Bandung Jawa Barat Indonesia}

	/* ===== POINTER =====
	- Kemampuan membuat reference ke lokasi data di memory yg sama, tanpa menduplikasi data yg sudah ada
	- Dengan kemampuan pointer, kita bisa membuat pass by reference
	- Untuk membuat sebuah variable dengan nilai pointer ke variable lain, kita bisa menggunakan operator
	  & diikuti dengan nama variablenya
	*/
	var address3 *Address = &address1
	address3.City = "Bogor"
	fmt.Println(address1) // {Bogor Jawa Barat Indonesia}
	fmt.Println(address3) // &{Bogor Jawa Barat Indonesia}
}

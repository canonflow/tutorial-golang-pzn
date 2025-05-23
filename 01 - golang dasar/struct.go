package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	/* ===== INTRO =====
	- Struct adalah template data yg digunakan untuk menggabungkan nol atau lebih tipe data lainnya dalam 1 kesatuan
	- Merepresentasikan data dalam program aplikasi yg kita buat
	- Data disimpan di dalam field / Atribut
	- Singkatnya, Struct adalah kumpulan dari field / Atribut
	*/
	var nathan Customer
	nathan.Name = "Nathan"
	nathan.Address = "Indonesia"
	nathan.Age = 21
	fmt.Println(nathan)      // {Nathan Indonesia 21}
	fmt.Println(nathan.Name) // Nathan

	// Struct Literals
	joko := Customer{
		Name:    "Joko",
		Address: "Indonesia",
		Age:     21,
	}
	fmt.Println(joko) // {Joko Indonesia 21}

	budi := Customer{"Budi", "Indonesia", 25}
	fmt.Println(budi) // {Budi Indonesia 25}
}

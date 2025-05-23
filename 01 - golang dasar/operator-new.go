package main

import "fmt"

type Home struct {
	City, Province, Country string
}

func main() {
	/* ===== INTRO =====
	- Sebelumnya, utk membuat pointer dengan menggunakan operator &
	- Golang jg memiliki function New yg bisa digunakan UNTUK MEMBUAT POINTER
	- Namun, function New hanya MENGEMBALIKAN POINTER ke DATA KOSONG, artinya TIDAK ADA DATA AWAL
	*/

	//var alamat1 *Home = &Home{} // Dapat diganti dengan operator new
	var alamat1 = new(Home)
	var alamat2 *Home = alamat1
	alamat2.Country = "Indonesia"

	fmt.Println(alamat1) // &{  Indonesia}
	fmt.Println(alamat2) // &{  Indonesia}

}

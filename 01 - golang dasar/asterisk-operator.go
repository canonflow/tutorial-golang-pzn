package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	/* ===== INTRO =====
	- Saat kita mengubah variable pointer, maka yang berubah HANYA variable tsb
	- Semua variable yg mengacu ke data yg sama TIDAK AKAN BERUBAH
	- Jika kita ingin mengubah seluruh variable yg mengacu ke data tsb, kita bisa menggunakan operator *
	*/

	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}

	var address2 *Address = &address1
	address2.City = "Bandung"
	fmt.Println(address1) // {Bandung Jawa Barat Indonesia}
	fmt.Println(address2) // &{Bandung Jawa Barat Indonesia}

	// ===== HANYA MENGUBAH address2 SAJA =====
	//address2 = &Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	//fmt.Println(address1) // {Bandung Jawa Barat Indonesia}
	//fmt.Println(address2) // &{Jakarta DKI Jakarta Indonesia}

	// ===== MENGUBAH POINTER SEMUA VARIABLE =====
	*address2 = Address{"Kediri", "Jawa Timur", "Indonesia"}
	fmt.Println(address1) // {Kediri Jawa Timur Indonesia}
	fmt.Println(address2) // &{Kediri Jawa Timur Indonesia}
}

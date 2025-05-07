package main

import "fmt"

func main() {
	// SLICE
	/* ===== DETAIL =====
	3 data penting pada Slice:
		1. pointer -> petunjuk data pertama di array para slice
		2. length -> panjang dari slice
		3. capacity -> kapasitas dari slice, dimana length tidak boleh melebihi dari capacity

	Membuat slice dari array
		1. array[low:high] -> membuat slice dimulai dari index low sampai index sebelum high
		2. array[low:] -> membuat slice dimulai dari index low sampai index terakhir di array
		3. array[:high] -> membuat slice dimulai dari index 0 sampai index sebelum high
		4. array[:] -> membuat slice dari index 0 sampai index terakhir di array
	*/

	names := [...]string{
		"Nathan", "Garzya", "Santoso",
		"Joko", "Budi", "Nugraha",
	}

	slice1 := names[4:6]
	fmt.Println(slice1[0]) // Budi
	fmt.Println(slice1[1]) // Nugraha

	slice2 := names[:3]
	fmt.Println(slice2) // [Nathan Garzya Santoso]

	slice3 := names[3:]
	fmt.Println(slice3) // [Joko Budi Nugraha]

	var slice4 []string = names[:]
	fmt.Println(slice4) // [Nathan Garzya Santoso Joko Budi Nugraha]

	/* ===== FUNCTION =====
	1. len(slice) -> mendapatkan panjang
	2. cap(slice) -> mendapatkan kapasitas
	3. append(slice, data) -> membuat slice baru dengan menambahkan data ke posisi terakhir slice, jika kapasitas
								penuh, maka akan membuat array baru
	4. make([]TypeData, length, capacity) -> membuat slice baru
	5. copy(destination, source) -> menyalin slice dari source ke destination
	*/

	// ===== Append =====
	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	daySlice1 := days[5:]
	// daySlice1 menggunakan array dari arra days, maka dari itu ketika kita mengganti value dari daySlice1
	// sebenarnya kita mengganti nilai dari array days
	daySlice1[0] = "Sabtu baru"
	daySlice1[1] = "Minggu baru"

	fmt.Println(days) // [Senin Selasa Rabu Kamis Jumat Sabtu baru Minggu baru]

	// Karena kita tahu bahwa kapasitas dari slice1 sudah penuh, maka ketika kita menambahkan data baru
	// sistem akan membuat array baru
	daySlice2 := append(slice1, "Libur baru")
	daySlice2[0] = "Ups"
	fmt.Println(daySlice2) // [Ups Nugraha Libur baru]
	fmt.Println(days)      // [Senin Selasa Rabu Kamis Jumat Sabtu baru Minggu baru]

	// ===== MAKE =====
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Nathan"
	newSlice[1] = "Garzya"
	//newSlice[2] = "Joko" // error, harusnya menggunakan append

	fmt.Println(len(newSlice)) // 2
	fmt.Println(cap(newSlice)) // 5

	newSlice2 := append(newSlice, "Joko")
	fmt.Println(newSlice2)      // [Nathan Garzya Joko]
	fmt.Println(len(newSlice2)) // 3
	fmt.Println(cap(newSlice))  // 5

	newSlice2[0] = "UBAH"
	fmt.Println(newSlice2) // [UBAH Garzya Joko]
	fmt.Println(newSlice)  // [UBAH Garzya]

	// ===== COPY =====
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)
	fmt.Println(fromSlice) // [Senin Selasa Rabu Kamis Jumat Sabtu baru Minggu baru]
	fmt.Println(toSlice)   // [Senin Selasa Rabu Kamis Jumat Sabtu baru Minggu baru]

	// ARRAY VS SLICE
	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(iniArray) // [1 2 3 4 5]
	fmt.Println(iniSlice) // [1 2 3 4 5]
}

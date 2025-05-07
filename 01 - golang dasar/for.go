package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke", counter)
		counter++
	}

	/* ===== FOR WITH STATEMENT =====
	- init statement -> statement sebelum for dieksekusi
	- post statement -> statement yang akan dieksekusi di akhir tiap perulangan
	*/

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke", counter)
	}

	/* ===== FOR RANGE =====
	- For bisa digunakan untuk melakukan iterasi terhadap semua data collection
	- Data collection contohnya Array, Slice, dan Map
	*/
	names := []string{"Nathan", "Garzya", "Santoso"}
	// Manual
	//for i := 0; i < len(names); i++ {
	//	fmt.Println(names[i])
	//}

	for index, name := range names {
		fmt.Println(index, name)
	}

	// Kalo index-nya gk butuh, ganti dengan _ (underscore)
	for _, name := range names {
		fmt.Println(name)
	}
}

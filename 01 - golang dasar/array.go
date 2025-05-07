package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Nathan"
	names[1] = "Garzya"
	names[2] = "Santoso"

	fmt.Println(names[0]) // Nathan
	fmt.Println(names[1]) // Garzya
	fmt.Println(names[2]) // Santoso

	// ===== Membuat Array secara langsung =====
	var values = [3]int{
		90,
		80,
		95,
	}

	fmt.Println(values) // [90 80 95]

	// ===== Function Array =====
	fmt.Println(len(values)) // 3
	values[0] = 100
	fmt.Println(values[0]) // 100

	// ===== ADD-ONS =====
	// ... akan menghitung jumlah item yg ada di dalam array
	var scores = [...]int{
		100,
		90,
		80,
	}
	fmt.Println(scores) // [100 90 80]
}

package main

import "fmt"

func main() {
	// Break
	for i := 0; i < 10; i++ {
		if i == 5 {
			break // Akan menghentikan For loop
		}
		fmt.Println("Perulangan ke", i)
	}

	// Continue
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue // Kalo genap, lanjutkan perulangan selanjutnya
		}
		fmt.Println("Perulangan ke", i)
	}
}

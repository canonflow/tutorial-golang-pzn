package main

import "fmt"

func main() {
	// ===== CONSTANT =====
	// Kalo tidak dipakai, tidak masalah. Beda dengan variable biasa
	const firstName string = "Nathan"
	const lastName = "Santoso"

	// Error
	// firstName = "Garzya"

	// ===== MULTIPLE CONSTANTS =====
	const (
		street string = "Raya Tenggilis"
		hobby         = "Swimming"
	)
	fmt.Println(street, hobby)
}

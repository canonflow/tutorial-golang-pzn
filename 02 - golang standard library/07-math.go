package main

import (
	"fmt"
	"math"
)

func main() {
	/* ===== PACKAGE MATH =====
	- Round(float64): membulatkan ke atas / bawah, sesuai dengan yg paling dekat
	- Floor(float64): membulatkan ke bawah
	- Ceil(float64): membulatkan ke atas
	- Max(float64, float64): Mengembalikan nilai float64 terbesar
	- Min(float64, float64): Mengembalikan nilai float64 terkecil
	*/
	fmt.Println(math.Ceil(1.40))  // 2
	fmt.Println(math.Floor(1.60)) // 1
	fmt.Println(math.Round(1.50)) // 2
	fmt.Println(math.Round(2.50)) // 3
	fmt.Println(math.Max(10, 11)) // 11
	fmt.Println(math.Min(10, 11)) // 10
}

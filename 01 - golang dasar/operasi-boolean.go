package main

import "fmt"

func main() {
	var nilaiAkhir = 90
	var absensi = 81

	var lulusNilaiAkhir bool = nilaiAkhir > 80
	var lulusAbsensi bool = absensi > 80

	var lulus bool = lulusAbsensi && lulusNilaiAkhir
	fmt.Println(lulus) // true
}

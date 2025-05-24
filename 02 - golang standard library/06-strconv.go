package main

import (
	"fmt"
	"strconv"
)

func main() {
	/* ===== INTRO =====
	- Sblmnya, kita sudah belajar cara konversi tipe data misal dari int32 ke int64
	- Bagaimana jika kita btuh melakukan konversi yg tipe datanya berbeda? Misal dari int ke string, atau sebaliknya
	- Hal tsb bisa kita lakukan dengan bantuan package strconv (string conversion)
	- Beberapa Function:
		- Atoi(string): string ke int
		- Itoa(int): int ke string
		- ParseBool(string): string ke bool
		- ParseFloat(string): string ke float
		- ParseInt(string): string ke int64
		- FormatBool(bool): bool ke string
		- FormatFloat(float, ...): float64 ke string
		- FormatInt(int, base, bitSize): int ke string
	*/
	boolean, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println("Result:", boolean) // true
	}

	//integer, err := strconv.ParseInt("1000", 10, 64)
	integer, err := strconv.Atoi("1000")
	if err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println("Result:", integer) // 1000
	}

	binary := strconv.FormatInt(999, 2)
	fmt.Println("Binary 999:", binary) // 1111100111

	var stringInt string = strconv.Itoa(999)
	fmt.Println("String 999:", stringInt) // 999

}

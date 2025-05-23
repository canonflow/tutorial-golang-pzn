package main

import (
	"fmt"
	"golang-dasar/helper"
)

func main() {
	/* ===== INTRO =====
	- Secara standar, file golang hanya bisa mengakses file golang lainnya yg berada dalam package yg sama
	- Jika kita ingin mengakses file golang yg berada diluar package, maka kita bisa menggunakan Import
	*/
	result := helper.SayHello("Nathan")
	fmt.Println(result)             // Hello Nathan
	fmt.Println(helper.Application) // golang
	//fmt.Println(helper.version)              // gk isa
	//fmt.Println(helper.sayGoodBye("Nathan")) // gk isa
	helper.Contoh()
}

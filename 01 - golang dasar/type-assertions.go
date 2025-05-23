package main

import "fmt"

func random() interface{} {
	return true
}

func main() {
	/* ===== INTRO =====
	- Type Assertions merupakan kemampuan tipe data menjadi tipe data yang diinginkan
	- Fitur ini sering sekali digunakan ketika kita bertemu dengan data interface kosong (any)
	*/
	var result any = random()
	//resultString := result.(string)
	//fmt.Println(resultString) // "OK"

	//resultInt := result.(int) // panic
	//fmt.Println(resultInt)

	// ===== CHECK TIPE DATA =====
	switch value := result.(type) {
	case string:
		// Value di sini sudah menjadi String
		fmt.Println("String", value)
	case int:
		// Value di sini sudah menjadi Integer
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknown", value)
	}
}

package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	/* ===== INTRO =====
	- Di dalam bahasa lain, object yg blm diinisialisasi maka secara otomatis nilainya adalah null / nil
	- Berbeda dengan Golang, di Golang saat kita buat variabel dengan tipe data tertentu, maka secara otomatis
	  akan dibuatkan default valuenya
	- Namun, di Golang ada data NIL, yaitu data kosong
	- Nil sendiri hanya bisa digunakan di BEBERAPA TIPE DATA, seperti
		- Interface
		- Function
		- Map
		- Slice
		- Pointer
		- Channel
	*/
	data := NewMap("Nathan")
	if data == nil {
		fmt.Println("Data map masih kosong")
	} else {
		fmt.Println(data["name"])
	}

}

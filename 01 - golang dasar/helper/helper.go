package helper

import "fmt"

/* ===== INTRO PACKAGE =====
- Package adalah tempat yg bisa digunakan utk mengorganisir kode program yg kita buat di Golang
- Dengan menggunakan package, kita bisa merapikan kode program yg kita buat
- Package sendiri sebenarnya hanya direktori folder di sistem operasi kita
- Untuk menentukan Access Modifier, ckup dengan nama function / variable
- Jika nama diawali dengan huruf besar, artinya bisa diakses dari package lain
- Jika dimulai dengan huruf kecil, artinya tidak bisa diakses dair package lain
*/

var version = "1.0.0" // Tidak bisa diakses dari luar
var Application = "golang"

// Tidak bisa diakses dari luar package
func sayGoodBye(name string) string {
	return fmt.Sprintf("Goodbye, %s!", name)
}

func Contoh() {
	sayGoodBye("Garzya")
	fmt.Println(version)
}

func SayHello(name string) string {
	return "Hello " + name
}

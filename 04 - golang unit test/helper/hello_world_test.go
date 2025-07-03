package helper

import "testing"

/* ===== POLICIES =====
-- FILE TEST --
	- Membuat file utk unit test harus menggunakan akhiran _test (e.g. hello_world_test.go)

-- FUNCTION --
	- Membuat function harus diawali dengan nama Test dilanjutkan dengan nama function yg akan diuji
		- TestHelloWorld
	- Harus memiliki parameter (t testing.T) dan tidak mengembalikan return value
*/

/* ===== HOW TO RUN =====
- go test: menjalankan unit test kit
- go test -v: melihat lebih detail function test apa saja yang sudah dirunning
- go test -v -run TestNamaFunction: memilih function unit test mana yang ingin dirunning
- go test./...: menjalankan semua unit test dari TOP FOLDER module-nya
*/

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Nathan")
	if result != "Hello Nathan!" {
		// Unit test failed
		panic("Result is not 'Hello Nathan!'")
	}
}

func TestHelloWorldGarzya(t *testing.T) {
	result := HelloWorld("Garzya")
	if result != "Hello Garzya!" {
		// Unit test failed
		panic("Result is not 'Hello Garzya!'")
	}
}

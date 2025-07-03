package helper

import (
	"fmt"
	"testing"
)

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

/*
- t.Fail(): menggagalkan unit test, namun tetap MELANJUTKAN eksekusi unit test. Namun di akhir ketika selesai,
			maka unit test tersebut dianggap gagal
- t.FailNow():menggagalkan unit test saat ini juga, TANPA MELANJUTKAN eksekusi unit test
- t.Error(args...): melakukan log (print) error, namun setelah melakukan log error, dia akan secara otomatis
					memanggil t.Fail(), artinya eksekusi unit test akan tetap dijalankan sampai selesai
- t.Fatal(args...): mirip dengan t.Error(), hanya saja setelah melakukan log error, dia akan memanggil
					t.FailNow(), sehingga mengakibatkan eksekusi unit test berhenti
*/

func TestHelloWorldNathan(t *testing.T) {
	result := HelloWorld("Nathan")
	if result != "Hello Nathan!" {
		// Unit test failed
		//panic("Result is not 'Hello Nathan!'")
		//t.Fail()
		t.Error("Result must be 'Hello Nathan!', got: " + result)
	}

	// Tetap dijalankan
	fmt.Println("===== TestHelloWorldNathan DONE =====")
}

func TestHelloWorldGarzya(t *testing.T) {
	result := HelloWorld("Garzya")
	if result != "Hello Garzya!" {
		// Unit test failed
		//panic("Result is not 'Hello Garzya!'")
		//t.FailNow()
		t.Fatal("Result must be 'Hello Garzya!', got: " + result)
	}

	// Tidak akan dijalankan
	fmt.Println("===== TestHelloWorld DONE =====")
}

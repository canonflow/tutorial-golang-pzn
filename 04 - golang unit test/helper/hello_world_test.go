package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
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
- t.Skip(): Untuk membatalkan bukan menggagalkan unit test yg kita mau
*/

/* ===== ASSERTION =====
- assert: Jika gagal, menggunakan t.Fail()
- require: Jika gagal, menggunakan t.FailNow()
*/

/*
	===== BEFORE AFTER TEST =====
  - Jika terdapat function bernama TestMain, golang akan mengeksekusi function ini tiap kali akan menjalankan
    test di sebuah package
  - Dengan ini kita bisa mengatur Before dan After unit test sesuai dengan yang kita mau

-Ingat, function TestMain itu dieksekusi hanya sekali per Golang package, bukan per tiap function unit test
*/

/* ===== SUB TEST =====
- go test -run TestNamaFunction/NamaSubTest: menjalankan hanya salah satu sub test
- go test -run /NamaSubTest: menjalankan semua test dengan semua sub test di semua function
*/

/* ===== BENCHMARK =====
- go test -v -bench . : Menjalankan seluruh benchmark di module
- go test -v -run=NotMathUnitTest -bench . : Menjalankan benchmark tanpa unit test
	-> menjalankan function unit test yg tidak ada (dgn kata lain, tidak ada unit test yg dijalankan)
- go test -v -run=NotMatchUnitTest -bench=BenchamarkTest : Menjalankan tanpa unit test dengan benchmark tertentu
- go test -v -bench ../...
- go test -v -bench BenchmarkSub/NamaSubBenchmark: Menjalankan sub benchmark

----- CONTOH OUTPUT -----
goos: windows
goarch: amd64
pkg: golang-unit-test/helper
cpu: AMD Ryzen 7 5800H with Radeon Graphics
BenchmarkHelloWorld
BenchmarkHelloWorld-16                  72040485                16.79 ns/op
BenchmarkHelloWorldGarzya
BenchmarkHelloWorldGarzya-16            71042908                16.48 ns/op
*/

func BenchmarkSub(b *testing.B) {
	b.Run("Nathan", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Nathan")
		}
	})

	b.Run("Garzya", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Garzya")
		}
	})
	/*
		goos: windows
		goarch: amd64
		pkg: golang-unit-test/helper
		cpu: AMD Ryzen 7 5800H with Radeon Graphics
		BenchmarkSub
		BenchmarkSub/Nathan
		BenchmarkSub/Nathan-16          72196518                16.91 ns/op
		BenchmarkSub/Garzya
		BenchmarkSub/Garzya-16          71773768                16.73 ns/op
	*/
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Nathan")
	}
}

func BenchmarkHelloWorldGarzya(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Garzya")
	}
}

func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "HelloWorld(Nathan)",
			request:  "Nathan",
			expected: "Hello Nathan!",
		},
		{
			name:     "HelloWorld(Garzya)",
			request:  "Garzya",
			expected: "Hello Garzya!",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Nathan", func(t *testing.T) {
		result := HelloWorld("Nathan")
		assert.Equal(t, "Hello Nathan!", result, "Result must be 'Hello Nathan!'")
	})

	t.Run("Garzya", func(t *testing.T) {
		result := HelloWorld("Garzya")
		assert.Equal(t, "Hello Garzya!", result, "Result must be 'Hello Garzya!'")
	})

}

func TestMain(m *testing.M) {
	// Before
	fmt.Println("===== BEFORE UNIT TEST =====")

	m.Run()

	// After (tetap dijalankan walaupun ada unit test yang fail (FailNow, require)
	fmt.Println("===== AFTER UNIT TEST =====")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Can't run on Mac OS")
	}

	result := HelloWorld("Nathan")
	assert.Equal(t, "Hello Nathan!", result)
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Nathan")
	assert.Equal(t, "Hello Nathan!", result, "Result must be 'Hello Nathan!'")
	fmt.Println("===== TestHelloWorldAssert DONE =====")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Nathan")
	require.Equal(t, "Hello Nathan!", result, "Result must be 'Hello Nathan!'")
	fmt.Println("===== TestHelloWorldRequire DONE =====")
}

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

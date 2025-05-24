package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	/* ===== INTRO =====
	- Package container/ring adalah implementasi struktur data circular list
	- Circular list adalah struktur data ring, dimana di akhir element akan kembali ke element awal (HEAD)
	*/
	var data *ring.Ring = ring.New(5)
	// Mengisi data
	for i := 0; i < data.Len(); i++ {
		data.Value = "Value " + strconv.Itoa(i)
		data = data.Next()
	}

	// Menampilkan data
	data.Do(func(value any) {
		fmt.Println(value)
	})
}

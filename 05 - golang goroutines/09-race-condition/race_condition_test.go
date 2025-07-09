package _9_race_condition

import (
	"fmt"
	"testing"
	"time"
)

/* ===== INTRO =====
- Saat kita menggunakan Goroutine, kita tidak hanya berjalan secara conccurent, tetapi bisa parallel juga,
  karena bisa ada beberapa thread yang berjalan secara paralell
- Hal ini sangat BERBAHAYA ketika kita melakukan manipulasi data variabel YANG SAMA oleh BEBERAPA GOROUTINE
  SECARA BERSAMAAN
- Hal ini bisa menyebabkan masalah yang namanya RACE CONDITION
*/

func TestRaceCondition(t *testing.T) {
	x := 0

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				x = x + 1
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter: ", x) // Seharusnya 100000 (100rb)
	/*
		=== RUN   TestRaceCondition
		Counter:  98423
		--- PASS: TestRaceCondition (5.00s)
		PASS
	*/
}

package _16_sync_map

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/* ===== INTRO =====
- Golang memiliki sebuah struct bernama sync.Map
- Map ini mirip dengan Map di Golang, namun yang membedakan, Map ini AMAN untuk penggunaan CONCURRENT menggunakan goroutine.
- Ada beberapa function yang bisa kita gunakan di Map:
	- Store(key, val): untuk menyimpan data ke map
	- Load(key): untuk mengambil data dari map menggunakan key
	- Delete(key): untuk menghapus data dari map menggunakan key
	- Range(func(key, val)): untuk melakukan iterasi seluruh data di map
*/

func TestSyncMap(t *testing.T) {
	var data sync.Map
	addToMap := func(value int) {
		data.Store(value, value)
	}

	for i := 0; i < 100; i++ {
		go addToMap(i)
	}

	time.Sleep(3 * time.Second)
	data.Range(func(key, value interface{}) bool {
		fmt.Println(key, ":", value)
		return true
	})

	/*
		=== RUN   TestSyncMap
		24 : 24
		25 : 25
		53 : 53
		41 : 41
		45 : 45
		94 : 94
		2 : 2
		6 : 6
		15 : 15
		22 : 22
		37 : 37
		54 : 54
		91 : 91
		0 : 0
		72 : 72
		95 : 95
		88 : 88
		17 : 17
		1 : 1
		39 : 39
		55 : 55
		59 : 59
		67 : 67
		61 : 61
		70 : 70
		4 : 4
		32 : 32
		49 : 49
		12 : 12
		18 : 18
		60 : 60
		83 : 83
		84 : 84
		78 : 78
		81 : 81
		77 : 77
		14 : 14
		33 : 33
		36 : 36
		52 : 52
		89 : 89
		92 : 92
		90 : 90
		27 : 27
		26 : 26
		66 : 66
		43 : 43
		8 : 8
		11 : 11
		31 : 31
		57 : 57
		64 : 64
		75 : 75
		96 : 96
		5 : 5
		23 : 23
		29 : 29
		71 : 71
		99 : 99
		85 : 85
		79 : 79
		13 : 13
		30 : 30
		51 : 51
		65 : 65
		68 : 68
		73 : 73
		74 : 74
		16 : 16
		63 : 63
		82 : 82
		87 : 87
		19 : 19
		7 : 7
		9 : 9
		10 : 10
		20 : 20
		38 : 38
		48 : 48
		46 : 46
		3 : 3
		76 : 76
		93 : 93
		98 : 98
		69 : 69
		35 : 35
		47 : 47
		42 : 42
		44 : 44
		58 : 58
		62 : 62
		86 : 86
		34 : 34
		97 : 97
		28 : 28
		40 : 40
		50 : 50
		56 : 56
		80 : 80
		21 : 21
		--- PASS: TestSyncMap (3.00s)
		PASS
	*/
}

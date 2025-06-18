package main

import (
	"fmt"
	"path"
	"path/filepath"
)

func main() {
	/* ===== INTRO =====
	- Package Path digunakan untuk memanipulasi data path seperti path di URL atau path di File System
	- Secara default Package Path menggunakan slash sebagai karakter path-nya, oleh karena itu cocok untuk data URL
	- Namun jika ingin menggunakan untuk memanipulasi path di File System, karena windows menggunakan back slashes, maka
	  khusus untuk File System, perlu menggunakan package path/filepath
	*/

	// ===== PATH =====
	fmt.Println("===== PATH =====")
	fmt.Println(path.Dir("hello/world.go"))             // hello
	fmt.Println(path.Base("hello/world.go"))            // world.go
	fmt.Println(path.Ext("hello/world.go"))             // .go
	fmt.Println(path.Join("hello", "world", "main.go")) // hello/world/main.go

	// ===== PATH/FILEPATH =====
	fmt.Println("\n===== FILE PATH =====")
	fmt.Println(filepath.Dir("hello/world.go"))             // hello
	fmt.Println(filepath.Base("hello/world.go"))            // world.go
	fmt.Println(filepath.Ext("hello/world.go"))             // .go
	fmt.Println(filepath.IsAbs("hello/world.go"))           // false
	fmt.Println(filepath.IsLocal("hello/world.go"))         // true
	fmt.Println(filepath.Join("hello", "world", "main.go")) // hello\world\main.go
}

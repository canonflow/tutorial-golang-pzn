package database

/*
		===== PACKAGE INITIALIZATION =====
	  - Saat kita membuat package, kita bisa membuat sbg function yg akan diakses ketika package kita diakses
	  - Ini cocok ketika contohnya jika package kita berisikan function2 utk berkomunikasi dengan database, kita
	    membuat function inisialisasi utk membuat koneksi ke database
	  - Untuk membuat function yg diakses secara otomatis ketika package diakses, kita cukup membuat function dengan nama
	    init
*/
var connection string

func init() {
	connection = "MySQL"
}

func GetDatabase() string {
	return connection
}

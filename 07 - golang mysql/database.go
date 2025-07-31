package main

import (
	"database/sql"
	"time"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3307)/belajar_golang_database?parseTime=true")

	if err != nil {
		panic(err)
	}
	//defer db.Close()

	db.SetMaxIdleConns(10)  // Minimal koneksi
	db.SetMaxOpenConns(100) // Maximal koneksi
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(1 * time.Hour)

	return db
}

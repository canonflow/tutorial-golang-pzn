package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestEmpty(t *testing.T) {
}

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3307)/belajar_golang_database")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	// Gunakan DB
}

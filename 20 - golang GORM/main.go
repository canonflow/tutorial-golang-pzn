package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func OpenConnection() *gorm.DB {
	dialect := mysql.Open("root:@tcp(localhost:3307)/belajar_golang_gorm?charset=utf8mb4&parseTime=True&loc=Local")

	db, err := gorm.Open(dialect, &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}

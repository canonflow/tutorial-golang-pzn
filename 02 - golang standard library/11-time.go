package main

import (
	"fmt"
	"time"
)

func main() {
	/* ===== INTRO =====
	- Package time adalah package yg berisikan fungsionalitas utk management waktu di golang
	- Beberapa function
		- Now(): mendapatkan waktu saat ini
		- Date(...): membuat waktu
		- Parse(layout, string): utk memparsing waktu dari string
	*/
	now := time.Now()
	fmt.Println(now.Local()) // 2025-05-24 15:22:48.5691848 +0700 +07

	//utc := now.UTC()
	utc := time.Date(2009, time.August, 17, 0, 0, 0, 0, time.UTC)
	fmt.Println(utc)         // 2009-08-17 00:00:00 +0000 UTC
	fmt.Println(utc.Local()) // 2009-08-17 07:00:00 +0700 +07

	formatter := "2006-01-02 15:04:05"
	value := "2020-10-10 10:10:10"
	//value := "ASAL"
	valueTime, err := time.Parse(formatter, value)
	if err == nil {
		fmt.Println(valueTime) // 2020-10-10 10:10:10 +0000 UTC
	} else {
		fmt.Println(err.Error())
	}

	fmt.Println(valueTime.Year())  // 2020
	fmt.Println(valueTime.Month()) // October
	fmt.Println(valueTime.Day())   // 10
	fmt.Println(valueTime.Hour())  // 10

	/* ===== DURATION =====
	- Saat menggunakan tipe data waktu, kadang kita perlu butuh data berupa durasi
	- Package tipe memiliki type Duration, yg sebenarnya adalah alias untuk int64
	- Namun terdapat banyak method yg bisa kita gunakan untuk memanipulasi Duration
	*/
	var duration1 time.Duration = time.Second * 1000
	var duration2 time.Duration = time.Millisecond * 10
	var duration3 time.Duration = duration1 - duration2
	fmt.Println(duration3)                  // 16m39.99s
	fmt.Printf("Duration: %d\n", duration3) // 999990000000 (format-nya nanosecond)
}

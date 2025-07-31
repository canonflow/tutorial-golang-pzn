# Golang Embed

---

## Pengenalan Embed Package
- Sejak Golang v1.16 terdapat package baru dengan nama Embed.
- Package embed adalah fitur baru untuk **mempermudah membaca isi file** pada **saat compile time secara otomatis** dimasukkan isi filenya dalam variable.
- Untuk melakukan embed file ke variable, kita bisa mengimport package embed terlebih dahulu
- Selanjutnya kita bisa tambahkan komentar `//go:embed` diikuti dengan **nama file**, **di atas variable yang dituju**
- Variable yang dituju tsb nanti **secara otomatis akan berisi** konten file yang kita inginkan secara otomatis ketika kode Golang dicompile
- Variable yang dituju tidak bisa disimpan di dalam function.

---

## Embed File ke String
- Embed file bisa kita lakukan ke variable dengan tipe data string.
- Secara **otomatis**, isi file akan **dibaca sbg text** dan **dimasukkan** ke variable tsb.
- 

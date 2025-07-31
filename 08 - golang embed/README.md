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

---

## Embed File ke []byte
- Selain ke tipe data String, embed file juga dapat dilakukan ke variable tipe data `[]byte`.
- Ini **cocok** sekali jika kita ingin melakukan embed file dalam **bentuk binary**, seperti **gambar** dan lain - lain.

---

## Embed Multiple Files
- Kadang ada kebutuhan kita ingin melakukan embed beberapa file sekaligus.
- Hal ini juga bisa dilakukan menggunakan embed package.
- Kita bisa menambahkan komentar `//go:embed` lebih dari 1 baris.
- Selain itu variable-nya bisa kita gunakan tipe data `embed.FS`.

---

## Path Matcher
- Selain manual 1 per 1, kita bisa menggunakan Path Matcher untul membaca multiple file yg kita inginkan.
- Ini sangat cocok ketika misal kita punya **pola jenis file** yang **kita inginkan** untuk dibaca.
- Caranya, kita perlu menggunakan path matcher seperti pada package function `path.Match`.



# Golang Generics

## Introduction
- Generic adalah kemampuan menambahkan parameter type saat membuat function.
- Berbeda dengan tipe data pada umumnya, generic dapat mengubah-ubah bentuk tipe data sesuai dengan yang kita mau.
- Fitur generics baru ada sejak Golang v1.18

### Manfaat
- **Pengecekan** ketika proses **kompilasi**
- **Tidak perlu manual** menggunakan **pengecekan tipe data** dan **konversi tipe** data.
- Memudahkan programmer membuat kode program yg generic sehingga bisa digunakan oleh berbagai tipe data.

---

## Type Parameter
- Untuk menandai sebuah function merupakan tipe generic, kita perlu menambahkan Type Parameter pada function tsb.
- Pembuatan Type Parameter menggunakna tanda **[] (kurung kotak)**, dimana **di dalam kurung kotak** tersebut, kita **tentukan nama** type Parameter-nya.
- Hampir sama dgn di bahasa pemrograman lain seperti Java, C#, dll, biasanya nama Type Parameter **hanya menggunakan 1 hurus**, misal T, K, V, dan lain-lain. Walaupun bisa saja lebih dari 1 huruf.

### Kode: Type Parameter
```go
func Length[T]() {
	
}
```

### Type Constraint
- Di bahasa pemrograman seperti Java, C#, Type Parameter biasanya tidak perlu kita tentukan tipe datanya, **berbeda** dgn di Golang.
- Di Golang, Type Parameter **wajib memiliki constraint**
- Type Constraint merupakan aturan yg digunakan untuk **menentukan tipe data yang diperbolehkan** pada Type Parameter.
- Contoh, jika kita ingin Type Parameter **bisa digunakan untuk semua tipe data**, kita bisa gunakan `interface{}` (kosong) sbg constraintnya.
- Type Constraint yg lebih detaul akan dibahas di materi **Type Sets**.

### Type Data any
- Di Golang 1.18, diperkenalkan **alias baru** bernama `any` untuk `interface{}`, ini bisa mempermudah kita ketika membuat Type Parameter dengan constraint `interface{}`, jadi kita cukup gunakan constraint `any`.

### Menggunakan Type Parameter
- Setelah kita buat Type Parameter di function, selanjutnya kita **bisa menggunakan** Type Parameter tsb sbg tipe data **di dalam function tsb**.
- Misalnya digunakan untuk return type atau function parameter.
- Kita cukup gunakan nama Type Parameternya saja.
- Type Parameter **hanya bisa digunakan di function saja**, tidak bisa digunakan di luar function.


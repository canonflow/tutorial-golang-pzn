# Golang Object Relational Mapping

## Pengenalan GORM

- ORM atau Object Relational Mapping adalah teknik untuk memetakan data dari database relational ke dalam object dalam pemrograman.
- Ketika kita belajar Golang Database, kita belajar pattern bernama **Repository**, yang digunakan sbg jembatan komunikasi ke Database.
- Saat meembuat Repository, kita membuat struct `Entity` sebagai **representasi** dari **table** di database.
- Hal itu sebenarnya sudah bisa dibilang sebuah `ORM`, namun masih dilakukan secara **manual**.

### Diagram Repository Pattern

![Diagram Repository Pattern](./assets/1.png)

### GORM

- GORM adalah salah satu library untuk implementasi ORM secara otomatis di Golang.
- Dengan menggunakan GORM, kita bisa fokus membuat pemetaan struct `Entity`, **tanpa harus memikirkan** detail dari implementasi `SQL` yang harus kita buat **untuk memanipulasi datanya**.
- GORM juga **mendukung relasi** antar `Entity` atau `Table`, baik **One to One**, **One to Many**, sampai **Many to Many**
- [https://gorm.io](https://gorm.io)

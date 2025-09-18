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

---

## Database Connection

- Untuk membuat koneksi ke database, kita bisa menggunakan `gorm.Open()`.
- **Tiap database** memiliki `config` masing-masing, kita bisa lihat semua config di database pada halaman: [https://gorm.io/docs/connecting_to_the_database.html](https://gorm.io/docs/connecting_to_the_database.html)

### Kode: Database Connection

```go

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

func OpenConnection() *gorm.DB {
    dialect := mysql.Open("root:@tcp(127.0.0.1:3306)/belajar_golang_gorm?charset=utfmb4&parseTime=True&loc=Local")

    db, err := gorm.Open(dialect, &gorm.Config{})
    if err != nil {
        panic(err)
    }

    return db
}
```

### Kode: Test Database Connection

```go
var db = OpenConnection()

func TestConnection(t *testing.T) {
    assert.NotNil(t, db)
}
```

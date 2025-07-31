# GOLANG MYSQL
___

## Database Pooling
- sql.DB di Golang sebenarnya bukanlah sebuah koneksi ke database
- Melainkan sebuah **pool** ke database, atau dikenal dengan konsep Database Pooling
- Di dalam sql.DB, Golang melakukan management koneksi ke database **secara otomatis**. Hal ini menjadikan kita tidak perlu melakukan management koneksi secara manual.

### Pengaturan Database Pooling
| Method | Keterangan|
| ------ | --------- |
| (DB) SetMaxIdleConns(number) | Pengaturan berapa jumlah **koneksi minimal** yang dibuat |
| (DB) SetMaxOpenConns(number) | Pengaturan berapa jumlah **koneksi maksimal** yang dibuat |
| (DB) SetConnMaxIdleTime(duration) | Pengaturan berapa **LAMA** koneksi yang sudah **TIDAK DIGUNAKAN** akan dihapus |
| (DB) SetConnMaxLifeTime(duration) | Pengaturan berapa lama koneksi **BOLEH DIGUNAKAN** |

---

## Eksekusi Perintah SQL
- Saat membuat aplikasi menggunakan database, sudah pasti kita ingin berkomunikasi dengan database menggunakan perintah SQL.
- Di Golang juga menyediakan function yang bisa kita gunakan untuk mengirim perintah SQL ke database menggunakan function **(DB) ExecContext(context, sql, params)**
- Ketika mengirim perintah SQL, kita butuh mengirimkan context, dan seperti yang sudah pernah kita pelajari di Course Golang Context, dengan context, kita bisa mengirim sinyal cancel jika kita ingin membatalkan pengiriman perintah SQL-nya.

---

## Query SQL
- Untuk operasi SQL yang **tidak membutuhkan hasil**, kita bisa menggunakan perintah **Exec**. Namun jika kita membutuhkan **result**, seperti **SELECT SQL**, kita bisa menggunakan function yang berbeda.
- Function untuk melakukan query ke database, bisa menggunakan function **(DB) QueryContext(context, sql, params)**.

### Rows
- Hasil Query Function adalah sebuah **data structs sql.Rows**
- Rows digunakan untuk melakukan iterasi terhadap hasil dari query.
- Kita bisa menggunakan function **(Rows) Next() (boolean)** untuk melakukan iterasi terhadap data hasisl query, jika return _**false**_, artinya sudah tidak ada data lagi di dalam struct.
- Untuk membaca tiap data, kita bisa menggunakan **(Rows) Scan(columns...)**
- Dan jangan lupa, setelah menggunakan **Rows**, jangan lupa untuk menutupnya menggunakan **(Rows) Close()**.

---

## Tipe Data Column
- Sebelumnya kita hanya membuat table dengan tipe data di kolom berupa **VARCHAR**
- Untuk **VARCHAR** di database, biasanya kita gunakan String di Golang.
- Bagaimana dengan tipe data yang lain?.
- Apa representasinya di Golang, misal tipe data timestamp, date, dan lainnya.

### Mapping Tipe Data
| Tipe Data Database | Tipe Data Golang |
| ------------------ | ---------------- |
| VARCHAR, CHAR | string |
| INT, BIGINT | int32, int64 |
| FLOAT, DOUBLE | float32, float64 |
| BOOLEAN | bool |
| DATE, DATETIME, TIME, TIMESTAMP | time.Time |

### Nullable Type
- Golang database **tidak mengerti** dengan tipe data NULL di database.
- Oleh karena itu, **khusus untuk kolom** yang bisa NULL di database, akan jadi masalah jika kita melakukan `Scan()` secara bulan - bulat menggunakan tipe data representasinya di Golang.

### Error Data Null
```
--- FAIL: TestQuerySqlComplex (0.00s)
panic: sql: Scan error on column index 2, name "email": converting NULL to string is unsupported [recovered]
	panic: sql: Scan error on column index 2, name "email": converting NULL to string is unsupported
```
- Konversi secara otomatis NULL tidak didukung oleh Driver MySQL Golang.
- Oleh karena itu, khusus tipe **kolom yang bisa NULL**, kita perlu **menggunakan tipe data** yang ada di dalam **package sql**.

### Tipe Data Nullable
| Tipe Data Golang | Tipe Data Nullable |
|------------------| ------------------ |
| string           | database/sql.NullString |
| bool             | database/sql.NullBool |
| float64          | database/sql.NullFloat64 |
| int32 | database/sql.NullInt32 |
| int64 | database/sql.NullInt64 |
| time.Time | database/sql.NullTime |

---

## SQL Injection
- Saat kita membuat aplikasi, kita tidak mungkin akan melakukan hardcode perintah SQL di kode Golang kita.
- Biasanya kita akan menggunakan input dari user, lalu membuat perintah SQL dari input user, dan mengirimnya menggunakan perintah SQL.
- SQL Injection adalah sebuah teknik yg menyalahgunakan sebuah celah keamanan yg terjadi dalam lapisan database sebuah aplikasi.
- Biasanya, SQL Injection dilakukan dengan mengirimkan input dari user dengan perintah yang salah, sehingga menyebabkan hasil SQL yang kita buat menjadi tidak valid.
- SQL Injection sangat berbahaya, jika sampai kita salah membuat SQL, bisa jadi data kita tidak aman.

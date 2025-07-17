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


## Eksekusi Perintah SQL
- Saat membuat aplikasi menggunakan database, sudah pasti kita ingin berkomunikasi dengan database menggunakan perintah SQL.
- Di Golang juga menyediakan function yang bisa kita gunakan untuk mengirim perintah SQL ke database menggunakan function **(DB) ExecContext(context, sql, params)**
- Ketika mengirim perintah SQL, kita butuh mengirimkan context, dan seperti yang sudah pernah kita pelajari di Course Golang Context, dengan context, kita bisa mengirim sinyal cancel jika kita ingin membatalkan pengiriman perintah SQL-nya.


## Query SQL
- Untuk operasi SQL yang **tidak membutuhkan hasil**, kita bisa menggunakan perintah **Exec**. Namun jika kita membutuhkan **result**, seperti **SELECT SQL**, kita bisa menggunakan function yang berbeda.
- Function untuk melakukan query ke database, bisa menggunakan function **(DB) QueryContext(context, sql, params)**.

### Rows
- Hasil Query Function adalah sebuah **data structs sql.Rows**
- Rows digunakan untuk melakukan iterasi terhadap hasil dari query.
- Kita bisa menggunakan function **(Rows) Next() (boolean)** untuk melakukan iterasi terhadap data hasisl query, jika return _**false**_, artinya sudah tidak ada data lagi di dalam struct.
- Untuk membaca tiap data, kita bisa menggunakan **(Rows) Scan(columns...)**
- Dan jangan lupa, setelah menggunakan **Rows**, jangan lupa untuk menutupnya menggunakan **(Rows) Close()**.
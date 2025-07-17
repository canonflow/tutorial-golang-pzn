# GOLANG MYSQL
___

## Database Pooling
- sql.DB di Golang sebenarnya bukanlah sebuah koneksi ke database
- Melainkan sebuah **pool** ke database, atau dikenal dengan konsep Database Pooling
- Di dalam sql.DB, Golang melakukan management koneksi ke database **secara otomatis**. Hal ini menjadikan kita tidak perlu melakukan management koneksi secara manual.

## Pengaturan Database Pooling
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


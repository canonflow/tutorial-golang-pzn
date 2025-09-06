# Golang Database Migration

## Pengenalan Database Migration

- Saat aplikasi berjalan, biasanya database sudah siap digunakan, artinya table, kolom, dan semua relasinya sudah dibuat di awal sebelum aplikasi berjalan.
- Apa yang terjadi ketika misal pd saat aplikasi sudah berjalan, kita perlu menambah fitur baru, lalu butuh mengubah struktur table di database?
- Biasanya kita akan mengubahnya di database langsung, lalu melakukan perubahan kode program.
- Hal ini mungkin terlihat sederhana, namun ketika sklanya sudah besar, dan anggota tim sudah banyak, maka perubahan langsung ke database bukanlah hal sederhana lagi.
- Kita harus bisa melakukan **tracking** apa saja **yang berubah**, dan memastikan semua anggota tim tahu perubahannya, sehingga bisa dilakukan hal yang sama di komputer masing2.

### Keuntungan Database Migration

- Database Migration adalah mekanisme untuk **melakukan tracking** perubahan struktur database, mulai dari awal dibuat sampai perubahaan terakhir yg dilakukan.
- Dengan menggunakan migration, semua tim member bisa melihat perubahan struktur database, dan bisa dengan mudah menjalankan perubahan tersebut di tiap komputer masing2.
- Selain itu, dengan adanya migration, kita bisa melakukan **review terlebih dahulu**, sblm menjalankan perubahan di database.

---

## Golang Migrate

- Golang Migrate adalah salah satu tool untuk Database Migration yg populer digunakan oleh programmer Golang.
- Golang migrate bisa diintegrasikan dengan aplikasi, atau dijalankan sbg aplikasi **standalnone**.
- Golang migrate mendukun banyak sekali database, seperti
  - MySQL
  - PostgreSQL
  - SQLITE
  - MongoDB
  - Cassandra
  - Dll

---

## Menginstall Golang Migrate

- Untuk menginstall Golang Migrate, sangat mudah. Kita bisa gunakan perintah berikut:
  - `go install -tags "database1,database2"
  - github.com/golang-migrate/migrate/v4/cmd/migrate@latest
- Sesuaikan dengand database yang ingin kita gunakan, bisa lebih dari 1 dengan cara menambahkan koma.

### Kode:

```shell
~ go install -tags 'postgres,mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

### Aplikasi Migrate

- Saat menginstall Golang Migrate, secara otomatis terdapat **executable file** di folder `$GOPATH/bin/` dengan nama **migrate**.
- **File migrate** tsb adalah **aplikasi Golang Migrate** yang akan kita gunakan untuk membuat Database Migration.

---

## Membuat Database Migration

- Untuk membuat database migration, kita bisa gunakan perintah:
  - `migrate create -ext sql -dir db/migration nama_file_migration`
- `-ext` adalah **file extension**, artinya kita membuat file `.sql.`.
- `-dir` adalah **folder tempat disimpan**.
- Usahakan tidak menggunakan spasi pada nama file migration.

### File Migration

- File migration akan **diawali dengan waktu** ketika kita membuat file migration, lalu diikuti dengan **nama migration** dan diakhir dengan **tipe migration**.
- misal `20220921103313_create_table_category.up.sql`
- Kenapa diawali dengan waktu? Agar file migration **selalu berurut** sesuai dengan waktu kita membuat file tsb.

---

## Migration Up

- File migration dengan akhiran **up** adalah file yang **harus kita isi** dengan **perubahan yang ingin kita tambahkan**.
- Misal, sekarang kita akan tambahkan table `categories`, sesuai dengan aplikasi RESTful API yang sudah kita buat.

---

## Migration Down

- Setiap file migration, selain up terdapat juga file **down**.
- File down adalah file yg berisikan kode untuk **mengembalikan perubahan yang kita lakukan di file up**.
- Kenapa diperlukan? Karena misal **terjadi masalah di aplikasi**, namin database migration **terlanjut dijalankan**, kita bisa melakukan `rollback` dengan cara **menjalankan file down**, karena berisikan kode unutk mengembalikan perubahan di file up.
- Pada kasus ini, misal kita akan menghapus lagi table `categories`.

---

## Membuat Database

- Sebelum menjalankan Database Migration, sekarang kita perlu membuat dulu database-nya.
- Hal ini karena pembuatan database tidak dilakukan di database migration, biasanya dilakukan manual di awal.
- Pada kasus ini, kita menggunakan database MySQL, dan kita perlu ubah kode koneksi database di aplikasi agar terhubung dengan database baru.

---

## Menjalankan Migration

- Kita dapat menjalankan database migration menggunakan perintah:
  - `migrate -database "koneksidatabase" -path folder up`
- `-database` harus berisikan **koneksi database**, misalnya untuk MySQL bisa menggunakan:
  - `mysql://user:password@tcp(host:port)/nama_database`
- Untuk database lain, bisa dilihat di halaman dokumentasinya: [https://github.com/golang-migrate/migrate#databases](https://github.com/golang-migrate/migrate#databases)
- `-path` adalah **lokasi folder file** database migration.
- `up` adalah perintah untuk menjalankan database migration dengan **mode up**.

---

## Migration State

- Saat kita sudah melakukan migration, lalu kita menambahkan file **migration baru**, apa yang terjadi jika kita menjalankan migration lagi?
- Golang Migrate akan menyimpan **state terakhir** kita menjalankan database migration, artinya **tidak akan dijalankan dari awal lagi**, melainkan dari file terakhir yang sukses di migrasi.
- Jadi kita tidak perlu takut file akan dijalankan lagi, jd tidak perlu hapus file migration-nya.
- Semua informasi **state** tersebut disimpan dalam table `schema_migrations`.

---

## Rollback Migration

- Pada waktu terjadi masalah pada aplikasi, yg menyebabkan kita harus melakukan **rollback** perubahan, apa yg harus dilakukan?
- Fitur itu sudah ada di Golang Migrate, jadi kita bisa menjalankan **mode down** untuk melakukan `rollback` dengan perintah:
  - `migrate -database "koneksidatabase" -path folder down

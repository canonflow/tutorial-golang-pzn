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

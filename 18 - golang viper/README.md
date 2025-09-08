# Golang Viper

## Pengenalan Golang Viper

- Setiap kali membuat aplikasi, kita tidak akan meng-harcode konfigurasi aplikasi di kode program kita. Misal konfigurasi koneksi ke database, atau sejenisnya.
- Hal ini dikarenakan, tiap tempat untuk menjalankan aplikasi yang kita buat mungkin akan **berbeda** konfigurasinya, misal di komputer kita, di server atau di komputer orang lain.
- Artinya, aplikasi yang kita buat, perlu memiliki kemampuan untuk mengambil data konfigurasi dari luar aplikasi, agar isi konfigurasi bisa diubah - ubah secara dinamis.

### Golang Viper

- Golang Viper salah satu library yg populer untuk melakukan manajemen konfigurasi, seperti:
  - JSON
  - YAML
  - env file
  - properties

---

## Membuat Viper

- Untuk membuat Viper, kita bisa menggunakan function `viper.New()`.
- Setelah membuat Viper, kita bisa menentukan **dari mana** kita akan mengambil konfigurasi.

### kode: Membuat Viper

```go
func TestViper(t *testing.T) {
    var config *viper.Viper = viper.New()
    assert.NotNil(t, config)
}
```

---

## JSON

- Viper bisa digunakan untuk membaca konfigurasi dari file JSON.

### Kode: Config JSON

```json
{
  "app": {
    "name": "golang-viper",
    "version": "1.0.0",
    "author": "Nathan Garzya"
  },
  "database": {
    "show_sql": true,
    "host": "localhost",
    "port": 3306
  }
}
```

### Kode: Membaca JSON

```go
func TestJson(t *testing.T) {
    config := viper.New()
    config.SetConfigName("config")
    config.SetConfigType("json")
    config.AddConfigPath(".")

    // Membaca
    err := config.ReadInConfig()
    assert.Nil(t, err)
}
```

### Mengambil Value

- Untuk **mengambil value** yang sudah dibaca dari file yang kita tentukan, kita bisa menggunakan method `GetXxx`, sesuai dengan tipe datanya, misal `GetString`, `GetInt`, `GetBool`, dan lain-lain.
- Jika kita ingin **mengakses nested object** pada `JSON`, kita bisa menggunakan `.`(titik), misal: `app.name`, `database.host`.

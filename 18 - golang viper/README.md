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

---

## YAML

- Selain JSON, Viper juga bisa digunakan untuk membaca file dengan formal `YAML`.
- Cara penggunaanya sama seperti ketika menggunakan `JSON`.

### Kode: Config YAML

```yaml
app:
  name: "golang-viper"
  version: "1.0.0"
  author: "Nathan Garzya"
database:
  show_sql: true
  host: "localhost"
  port: 3306
```

---

## ENV

- Viper juga bisa digunakan untuk membaca file dengan format `.env`.
- Di beberapa framework, kadang banyak yang menggunakan format ini.

### Kode: Config ENV

```env
APP_NAME=golang-viper
APP_VERSION=1.0.0
APP_AUTHOR=Nathan Garzya

DATABASE_HOST=localhost
DATABASE_PORT=3306
DATABASE_SHOW_SQL=true
```

---

## Environment Variable

- Kadang **saat menjalankan aplikasi**, kita **menyimpan konfigurasi** menggunakna `environment` variable yang terdapat di **sistem operasi yang kita gunakan**.
- Secara **default**, Viper **tidak akan membaca** data dari environment variable.
- Namun **jika mau**, kita bisa menggunakan method `AutomaticEnv()` **untuk membaca** dari environment variable.

---

## Fitur Lainnya

### File Config Lainnya

- Sebenarnya Viper bisa digunakan untuk membaca dari jenis file konfigurasi yang lain, misal:
- HCL (Hasicorp Configuration Language)
- Properties (Java Properties File)

### Remote Config

- Viper juga bisa digunakan untuk membaca konfigurasi dari **remote/server** yang terdapat di aplikasi:
  - Consul: [https://github.com/spf13/viper#consul](https://github.com/spf13/viper#consul)
  - Etcd: [https://github.com/spf13/viper#etcd](https://github.com/spf13/viper#etcd)

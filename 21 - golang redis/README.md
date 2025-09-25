# Golang Redis

---

## Redis

- Redis adalah salah satu In Memory database yang sangat populer, dan banyak digunakan
- Memahami Redis sekarang sudah menjadi salah satu hal yang sangat diperlukan, terutama untuk meningkatkan performa aplikasi yang kita buat
- Di kelas ini, kita akan belajar bagaimana berkomunikasi dengan Redis menggunakan Golang

### Golang Redis

- Redis sendiri sudah menyediakan library resmi yang bisa kita gunakan untuk di Golang
- Dengan begitu, berkomunikasi dengan Redis dari aplikasi Golang kita, akan sangat mudah [https://github.com/redis/go-redis](https://github.com/redis/go-redis)

---

## Client

- Hal pertama yang perlu kita lakukan saat ingin menggunakan Redis dari Golang adalah **membuat koneksi ke Redis.**
- Untuk membuat koneksi ke Redis, kita perlu membuat object `redis.Client`.
- Kita bisa menggunakan function `redis.NewClient(redis.Options)`.
- Kita bisa tentukan konfigurasi menggunakan `redis.Options`.

### Kode: Membuat Client

```go
var client = redis.NewClient(&redis.Options{
    Addr: "localhost:6379",
    DB: 0,
})

func TestConnection(t *testing.T) {
    assert.NotNil(t, client)

    err := client.Close()
    assert.Nil(t, err)
}
```

### Command

- Golang Redis sangat mudah digunakan, semua perintah Redis bisa kita lakukan menggunakan method yang terdapat di Client.
- Tiap `command` di Redis bisa kita lakukan di Client, dengan format **PascalCase**, misal `Ping()`, `Set()`, `Get()`, dan lainnya.
- Kita akan bahas tiap Command secara bertahap sesuai dengan struktur data yang akan kita gunakan.

### Kode: Ping

```go
var ctx = context.Background()
func TestPing(t *testing.T) {
    res, err := client.Ping(ctx).Result()

    assert.Nil(t, err)
    assert.Equal(t, "PONG", result)
}
```

---

## String

- Structur data yang **sering digunakan** di Redis adalah **String**.
- Command yang sering kita gunakan adalah menggunakan:
  - `Set()`
  - `SetEx()`
  - `Get()`
  - `MGet()`
  - dan lainnya

### Kode: String

```go
func TestString(t *testing.T) {
    client.SetEx(ctx, "name", "Nathan Garzya", time.Second * 5)

    result, err := client.Get(ctx, "name").Result()
    assert.Nil(t, err)
    assert.Equal(t, "Nathan Garzya", result)

    time.Sleep(time.Second * 5)
    result, err = client.Get(ctx, "name").Result()
    assert.NotNil(t, err)
}
```

---

## List

- Kita dapat menggunakan struktur data `List` di Golang Redis.

### Kode: List

```go
func TestList(t *testing.T) {
    client.RPush(ctx, "names", "nathan") // nathan
    client.RPush(ctx, "names", "garzya") // nathan, garzya
    client.RPush(ctx, "names", "santoso") // nathan, garzya, santoso

    assert.Equal(t, "nathan", client.LPop(ctx, "names").Val())
    assert.Equal(t, "garzya", client.LPop(ctx, "names").Val())
    assert.Equal(t, "santoso", client.LPop(ctx, "names").Val())

    client.Del(ctx, "names")
}
```

---

## Set

- Kita dapat menggunakan struktur data `Set` di Golang Redis.

### Kode: Set

```go
func TestSet(t *testing.T) {
    client.SAdd(ctx, "students", "nathan")
    client.SAdd(ctx, "students", "nathan")
    client.SAdd(ctx, "students", "garzya")
    client.SAdd(ctx, "students", "garzya")
    client.SAdd(ctx, "students", "santoso")
    client.SAdd(ctx, "students", "santoso")

    assert.Equal(t, int64(3), client.SCard(ctx, "students").Val())
    assert.Equal(t, []string{"nathan", "garzya", "santoso"}, client.SMembers(ctx, "students").Val())
}
```

---

## Sorted Set

- Kita juga dapat menggunakan struktur data `Sorted Set` di Golang Redis.
- **Default**, diurutkan berdasarkan terkecil `min, ..., max`

### Kode: Sorted Set

```go
func TestSortedSet(t *testing.T) {
    client.ZAdd(ctx, "scores", redis.Z{Score: 100, Member: "nathan"})
    client.ZAdd(ctx, "scores", redis.Z{Score: 85, Member: "garzya"})
    client.ZAdd(ctx, "scores", redis.Z{Score: 95, Member: "santoso"})

    assert.Equal(t, []string{"garzya", "santoso", "nathan"}, client.ZRange(ctx, "scores", 0, 2).Val())
    assert.Equal(t, "nathan", client.ZPopMax(ctx, "scores").Val()[0].Member)
    assert.Equal(t, "santoso", client.ZPopMax(ctx, "scores").Val()[0].Member)
    assert.Equal(t, "garzya", client.ZPopMax(ctx, "scores").Val()[0].Member)
}
```

---

## Hash

- Kita dapat menggunakan struktur data `Hash` di Golang Redis.

## Kode: Hash

```go
func TestHash(t *testing.T) {
    client.HSet(ctx, "user:1", "id", "1")
    client.HSet(ctx, "user:1", "name", "nathan")
    client.HSet(ctx, "user:1", "email", "nathan@example.com")

    user := client.HGetAll(ctx, "user:1").Val()
    assert.Equal(t, "1", user["id"])
    assert.Equal(t, "nathan", user["name"])
    assert.Equal(t, "nathan@example.com", user["email"])

    client.Del(ctx, "user:1")
}
```

---

## Geo Point

- Kita dapat menggunakan struktur data `Geo Point` di Golang Redis.

### Kode: Menambahkan Geo Point

```go
func TestGeoPoint(t *testing.T) {
    client.GeoAdd(ctx, "sellers", &redis.GeoLocation{
        Name: "Toko A",
        Longitude: 106.822702,
        Latitude: -6.177590,
    })
    client.GeoAdd(ctx, "sellers", &redis.GeoLocation{
        Name: "Toko B",
        Longitude: 106.820889,
        Latitude: -6.174964,
    })
}
```

### Kode: Mencari Geo Point

```go
func TestGeoPoint(t *testing.T) {
    // ......

    assert.Equal(t, 0.3543, client.GeoDist(ctx, "sellers", "Toko A", "Toko B", "km").Val())
    fmt.Println(client.GeoDist(ctx, "sellers", "Toko A", "Toko B", "km").Val())

    sellers := client.GeoSearch(ctx, "sellers", &redis.GeoSearchQuery{
        Longitude: 106.821825,
        Latitude: -6.175105,
        Radius: 5,
        RadiusUnit: "km",
    }).Val()

    assert.Equal(t, []string{"Toko A", "Toko B"}, sellers)
}
```

---

## Hyper Log Log

- Kita dapat menggunakan struktur data `Hyper Log Log` di Golang Redis.

### Kode: Hype Log Log

```go
func TestHyperLogLog(t *testing.T) {
    client.PFAdd(ctx, "visitor", "nathan", "garzya", "santoso")
    client.PFAdd(ctx, "visitor", "nathan", "canon", "flow")
    client.PFAdd(ctx, "visitor", "canon", "flow", "joko")

    assert.Equal(t, int64(6), client.PFCount(ctx, "visitors").Val())
}
```

---

## Pipeline

- Di Kelas Redis, kita pernah belajar tentang `pipeline`, dimana kita bisa **mengirim beberapa perintah** secara langsung **tanpa harus menunggu balasan satu per satu** dari Redis.
- Hal ini juga bisa dilakukan menggunakan Golang Redis menggunakan `Client.Pipelined(callback)`.
- Di dalam `callback`, kita bisa **melakukan semua command** yang akan **dijalankan dalam pipeline**.

### Kode: Pipeline

```go
func TestPipeline(t *testing.T) {
    client.Pipelined(ctx, func(pipeliner redis.Pipeliner) error {
        pipeliner.SetEx(ctx, "name", "Nathan", time.Second * 5)
        pipeliner.SetEx(ctx, "address", "Indonesia", time.Second * 5)
        return nil
    })

    assert.Equal(t, "Nathan", client.Get(ctx, "name").Val())
    assert.Equal(t, "Indonesia", client.Get(ctx, "address").Val())
}
```

---

## Transaction

- Kita tahu bahwa menggunakan Redis bisa melakukan transaction menggunakan perintah `MULTI` dan `COMMIT`. Namun **harus dalam koneksi yang sama**.
- Karena Golang Redis melakukan maintain **connection pool** secara **internal**, jadi kita **tidak bisa dengan mudah** menggunakan `MULTI` dan `COMMIT` menggunakan `client.Redis`.
- Kita harus menggunakan function `TxPipelined()`, dimana di dalamnya kita bisa membuat `callback` function **yang berisi command-command** yang akan dijalankan **dalam transaction**.

### Kode: Transaction

```go
func TestTransaction(t *testing.T) {
    client.TxPipelined(ctx, func(pipeliner redis.Pipeliner) error {
        pipeliner.SetEx(ctx, "name", "Nathan", time.Second * 5)
        pipeliner.SetEx(ctx, "address", "Surabaya", time.Second * 5)
        return nil
    })

    assert.Equal(t, "Nathan", client.Get(ctx, "name").Val())
    assert.Equal(t, "Surabaya", client.Get(ctx, "address").Val())
}
```

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

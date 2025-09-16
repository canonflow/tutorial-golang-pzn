# Golang Fiber

---

## Pengenalan Golang Fiber

### Web Framework

- Saat kita belajar kelas Golang Web, kita sudah belajar bagaimana cara membuat aplikasi web menggunakan Golang
- Dengan tambahan HTTPRouter, sebenarnya sudah cukup untuk membuat Web menggunakan Golang
- Namun, terkadang masih banyak beberapa hal yang harus dilakukan secara manual menggunakan bawaan dari Golang HTTP, sehingga kita membutuhkan Web Framework untuk mempermudah saat membuat aplikasi web menggunakan Golang
- Terutama Golang sangat populer untuk membuat RESTful API

### Fiber

- Sebenarnya ada banyak sekali Web Framework yang bisa kita gunakan untuk memudahkan kita membuat aplikasi Web dengan Golang
- Pada kelas ini, kita akan membahas salah satu Web Framework yang populer di Golang, yaitu Fiber
- Fiber adalah Web Framework untuk Golang yang terinspirasi dari ExpressJS, oleh karena itu cara penggunaannya sangat mudah dan sederhana
- Hal ini otomatis bisa mempermudah kita ketika membuat aplikasi Web / RESTful API menggunakan Golang
- https://gofiber.io/

---

## Fiber App

- Saat menggunakan Fiber, hal pertama yang perlu kita buat adalah `fiber.App`.
- Untuk mmebuatnya kita bisa menggunakan `fiber.New(fiber.Config)` yang menghasilkan pointer `*fiber.App`.
- `fiber.App` adalah representasi dari aplikasi **Web Fiber**.
- Setelah membuat `fiber.App`, selanjutnya untuk menjalankan aplikasi web-nya, kita bisa menggunakan method `Listen(address)`.

### Kode: Fiber App

```go
func main() {
    app := fiber.New()

    err := app.Listen("localhost:3000")

    if err != nil {
        panic(err)
    }
}
```

---

## Configuration

- Saat kita membuat `fiber.App` menggunakan `fiber.New()` terdapat parameter `fiber.Config` yang bisa kita gunakan.
- Ada banyak sekali konfigurasi yang bisa kita ubah, dan kita akan bahas secara bertahap.
- Contoh yang bisa kita gunakan adalah mengubah konfigurasi _tiemout_.

### Kode: Configuration

```go
func main() {
    app := fiber.New(fiber.Config{
        IdleTimeout: time.Second * 5,
        ReadTimeout: time.Second * 5,
        WriteTimeout: time.Second * 5,
    })

    err := app.Listen("localhost:3000")

    if err != nil {
        panic(err)
    }
}
```

---

## Routing

- Saat kita menggunakan Web Framework, pertama yang kita tanyakan adalah bagaimana cara membuat **endpoint-nya**.
- Di Fiber, untuk membuat **Routing**, sudah disediakan semua **Method** di `fiber.App` yang sesuai dengan HTTP Method.
- Parameter membutuhkan 2, yaitu `path` dan `fiber.Handler`.

### Kode: Routing Method

```go
func (app *App) Get(path string, handlers ...Handler) Router
func (app *App) Head(path string, handlers ...Handler) Router
func (app *App) Post(path string, handlers ...Handler) Router
func (app *App) Put(path string, handlers ...Handler) Router
func (app *App) Delete(path string, handlers ...Handler) Router
func (app *App) Connect(path string, handlers ...Handler) Router
func (app *App) Options(path string, handlers ...Handler) Router
func (app *App) Trace(path string, handlers ...Handler) Router
func (app *App) Patch(path string, handlers ...Handler) Router
```

### Kode: Hello World

```go
func TestRoutingHelloWorld(t *testing.T) {
    app := fiber.New()
    app.Get("/", func(ctx *fiber.Ctx) error {
        return ctx.SendString("Hello World")
    })

    request := httptest.NewRequest("GET", "/", nil)
    response, err := app.Test(request)

    // Test
    assert.Nil(t, err)
    assert.Equal(t, 200, response.StatusCode)

    bytes, err := io.ReadAll(response.Body)
    assert.Nil(t, err)
    assert.Equal(t, "Hello World", string(bytes))
}
```

---

## Ctx

- Saat kita membuat `Handler` di Fiber Routing, kita hanya cukup menggunakan parameter `fiber.Ctx`.
- **Ctx** ini merupakan representasi dari **Request** dan **Response** di Fiber.
- Oleh karena itu, kita **bisa mendapatkan informasi HTTP Request**, dan juga **bisa membuat HTTP Response** menggunakan fiber.Ctx.
- Untuk detail HTTP Request dan HTTP Response, akan kita bahas lebih detail di materi-materi selanjutnya.

### Kode: Ctx

```go
var app = fiber.New()

func TestCtx(t *testing.T) {
    app.Get("/hello", func(ctx *fiber.Ctx) error {
        name := ctx.Query("name", "Guest")
        return ctx.SendString("Hello", name)
    })

    request := httptest.NewRequest("GET", "/hello?name=Nathan", nil)
    response, err := app.Test(request)

    // Testing
    assert.Nil(t, err)
    bytes, err := io.ReadAll(response.Body)
    assert.Nil(t, err)
    assert.Equal(t, "Hello Nathan", string(bytes))
}
```

---

## HTTP Request

- Representasi dari HTTP Request di Fiber adalah `Ctx`.
- Untuk mengambil informasi dari HTTP Request, kita bisa menggunakan banyak sekali Method yang terdapat di Ctx.
- Kita bisa baca seluruh method yg tersedia di `Ctx` di halaman dokumentasinya.
- [https://pkg.go.dev/github.com/gofiber/fiber/v2#Ctx](https://pkg.go.dev/github.com/gofiber/fiber/v2#Ctx).

### Kode: HTTP Request

```go
func TestHTTPRequest(t *testing.T) {
    app.Get("/request", func(ctx *fiber.Ctx) error {
        first := ctx.Get("firstname") // Header
        last := ctx.Cookies("lastname") // Cookie
        return ctx.SendString("Hello", first, last)
    })

    request := httptest.NewRequest("GET", "/request", nil)
    request.Header.Set("firstname", "Nathan")
    request.AddCookie(&http.Cookie{Name: "lastname", value:"Santoso"})
    response, err := app.Test(request)

    assert.Nil(t, err)

    bytes, err := io.ReadAll(response.Body)
    assert.Nil(t, err)
    assert.Equal(t, "Hello Nathan Santoso", string(bytes))
}
```

---

## Route Parameter

- Fiber mendukung Route Parameter, dimana kita bisa menambahkan parameter di Path URL.
- Ini sangat cocok ketika kita membuat RESTful API, yang butuh data dikirim via Path URL.
- Saat membuat Route Parameter, kita perlu memberi **nama** dan di `Ctx`, kita **bisa mengambil seluruh data** menggunakan method `AllParams()`, atau menggunakan method `Params(nama)`.

### Kode: Route Parameter
```go
func TestRouteParam(t *testing.T) {
    app.Get("/users/:userid/orders/:orderId", func(ctx *fiber.Ctx) error {
        userId := ctx.Params("userId")
        orderId := ctx.Params("orderId")

        // return ctx.
    })

    request := httptest.NewRequest("GET", "/request", nil)
    request.Header.Set("firstname", "Nathan")
    request.AddCookie(&http.Cookie{Name: "lastname", value:"Santoso"})
    response, err := app.Test(request)

    assert.Nil(t, err)

    bytes, err := io.ReadAll(response.Body)
    assert.Nil(t, err)
    assert.Equal(t, "Hello Nathan Santoso", string(bytes))
}

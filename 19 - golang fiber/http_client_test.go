package main

import (
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

/* ===== INTRO =====
- Saat kita membuat aplikasi, terkadang kita jg sering menggunakan HTTP Client untuk mengambil atau mengirim data
  ke API / RESTful lainnya.
- Fiber menyediakan fitur untuk HTTP Client, sehingga kita bisa menggunakan Fiber sbg HTTP Client.
- Kita bisa menggunakan fiber.Client

===== CARA MEMBUAT CLIENT =====
- Fiber menyediakan Client Pool, agar ketika kita menggunakan Client dan selesai menggunakannya, kita bisa menggunakan ulang Client.
- Pool-nya dimanage oleh Fiber, jadi kita cukup ambil ketika membutuhkan menggunakan fiber.AcquireClient()
- Setelah selesai menggunakannya, kita harus mengembalikan Client ke Pool menggunakan fiber.ReleaseClient(client)

===== HTTP METHOD =====
- Saat menggunakan HTTP Client, kita bisa menggunakan method yang sesuai dengan HTTP Method yang mau kita gunakan.
- Misal, GET, POST, PUT, DELETE, PATCH, HEAD, OPTIONS
- Cara penggunaannya sangat mudah, hampir mirip dengan Ctx di Fiber.
*/

func TestClient(t *testing.T) {
	client := fiber.AcquireClient()

	agent := client.Get("https://example.com/")
	status, response, errors := agent.String()
	assert.Nil(t, errors)
	assert.Equal(t, 200, status)
	assert.Contains(t, response, "Example Domain")

	fiber.ReleaseClient(client)

	/*
		=== RUN   TestClient
		--- PASS: TestClient (0.84s)
		PASS
	*/
}

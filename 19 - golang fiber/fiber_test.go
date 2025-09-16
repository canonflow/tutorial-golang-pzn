package main

import (
	"bytes"
	_ "embed"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestRoutingHello(t *testing.T) {
	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello World!")
	})

	request := httptest.NewRequest("GET", "/", nil)
	response, err := app.Test(request)

	assert.Nil(t, err)

	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello World!", string(bytes))
	/*
		=== RUN   TestRoutingHello
		--- PASS: TestRoutingHello (0.00s)
		PASS
	*/
}

var app = fiber.New()

func TestCtx(t *testing.T) {
	app.Get("/", func(ctx *fiber.Ctx) error {
		name := ctx.Query("name", "Guest")
		return ctx.SendString("Hello " + name)
	})

	request := httptest.NewRequest("GET", "/?name=Nathan", nil)

	response, err := app.Test(request)

	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello Nathan", string(bytes))
	/*
		=== RUN   TestCtx
		--- PASS: TestCtx (0.00s)
		PASS
	*/
}

func TestHttpRequest(t *testing.T) {
	app.Get("/request", func(ctx *fiber.Ctx) error {
		first := ctx.Get("firstname")   // Header
		last := ctx.Cookies("lastname") // Cookie
		return ctx.SendString("Hello " + first + " " + last)
	})

	request := httptest.NewRequest("GET", "/request", nil)
	request.Header.Set("firstname", "Nathan")
	request.AddCookie(&http.Cookie{Name: "lastname", Value: "Garzya"})

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello Nathan Garzya", string(bytes))

	/*
		=== RUN   TestHttpRequest
		--- PASS: TestHttpRequest (0.00s)
		PASS
	*/
}

func TestRouteParameter(t *testing.T) {
	app.Get("/users/:userId/orders/:orderId", func(ctx *fiber.Ctx) error {
		userId := ctx.Params("userId")
		orderId := ctx.Params("orderId")

		return ctx.SendString("Get Order " + orderId + " From User " + userId)
	})
	request := httptest.NewRequest("GET", "/users/1/orders/2", nil)
	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Get Order 2 From User 1", string(bytes))

	/*
		=== RUN   TestRouteParameter
		--- PASS: TestRouteParameter (0.00s)
		PASS
	*/
}

func TestRequestForm(t *testing.T) {
	app.Post("/hello", func(c *fiber.Ctx) error {
		name := c.FormValue("name")

		return c.SendString("Hello " + name)
	})

	body := strings.NewReader("name=Nathan")
	request := httptest.NewRequest("POST", "/hello", body)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := app.Test(request)

	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)

	assert.Nil(t, err)
	assert.Equal(t, "Hello Nathan", string(bytes))

	/*
		=== RUN   TestRequestForm
		--- PASS: TestRequestForm (0.00s)
		PASS
	*/
}

//go:embed source/contoh.txt
var contohFile []byte

func TestMultipartForm(t *testing.T) {
	app.Post("/upload", func(c *fiber.Ctx) error {
		file, err := c.FormFile("file") // Get the file
		if err != nil {
			return err
		}

		err = c.SaveFile(file, "./target/"+file.Filename) // Save the file
		if err != nil {
			return err
		}

		return c.SendString("Upload Success")
	})

	// Setup Body
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)                    // Init the writer -> put the content into body
	file, _ := writer.CreateFormFile("file", "contoh.txt") // Create the Form File
	file.Write(contohFile)                                 // Put the file into Form File
	writer.Close()

	request := httptest.NewRequest("POST", "/upload", body)
	request.Header.Set("Content-Type", writer.FormDataContentType())

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Upload Success", string(bytes))

	/*
		=== RUN   TestMultipartForm
		--- PASS: TestMultipartForm (0.00s)
		PASS
	*/
}

package main

import (
	"io"
	"net/http/httptest"
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

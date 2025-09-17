package main

import (
	"errors"
	"fmt"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestRoutingGroup(t *testing.T) {
	helloWorld := func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	}

	api := app.Group("/api")
	api.Get("/hello", helloWorld)
	api.Get("/world", helloWorld)

	web := app.Group("/web")
	web.Get("/hello", helloWorld)
	web.Get("/world", helloWorld)

	request := httptest.NewRequest("GET", "/api/hello", nil)
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello World", string(bytes))

	/*
		=== RUN   TestRoutingGroup
		--- PASS: TestRoutingGroup (0.00s)
		PASS
	*/
}

func TestStatic(t *testing.T) {
	app.Static("/public", "./source")

	request := httptest.NewRequest("GET", "/public/contoh.txt", nil)
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "INI CONTOH FILE YANG AKAN DIUPLOAD YA!!", string(bytes))

	/*
		=== RUN   TestStatic
		--- PASS: TestStatic (0.04s)
		PASS
	*/
}

func TestErrorHandling(t *testing.T) {
	app.Get("/error", func(c *fiber.Ctx) error {
		return errors.New("UUPSS")
	})

	request := httptest.NewRequest("GET", "/error", nil)

	response, err := app.Test(request)

	assert.Nil(t, err)
	assert.Equal(t, 500, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Error : UUPSS", string(bytes))

	/*
		=== RUN   TestErrorHandling
		--- PASS: TestErrorHandling (0.00s)
		PASS
	*/
}

func TestView(t *testing.T) {
	app.Get("/view", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"title":   "Hello Title",
			"header":  "Hello Header",
			"content": "Hello Content",
		})
	})

	request := httptest.NewRequest("GET", "/view", nil)
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	fmt.Println(string(bytes))
	assert.Nil(t, err)
	assert.Contains(t, string(bytes), "Hello Title")
	assert.Contains(t, string(bytes), "Hello Header")
	assert.Contains(t, string(bytes), "Hello Content")

	/*
		=== RUN   TestView
		<!doctype html>
		<html lang="en">
		<head>
		    <meta charset="UTF-8">
		    <title>Hello Title</title>
		</head>
		<body>
		<h1>Hello Header</h1>
		<p>Hello Content</p>
		</body>
		</html>
		--- PASS: TestView (0.00s)
		PASS
	*/
}

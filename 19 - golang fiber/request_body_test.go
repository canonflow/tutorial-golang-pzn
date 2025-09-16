package main

import (
	"encoding/json"
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func TestRequestBody(t *testing.T) {
	app.Post("/login", func(c *fiber.Ctx) error {
		body := c.Body()
		request := new(LoginRequest)

		err := json.Unmarshal(body, request) // Parse the JSON into Struct (Unmarshall)
		if err != nil {
			return err
		}

		return c.SendString("Login Success " + request.Username)
	})

	body := strings.NewReader(`{"username": "Nathan", "password": "Garzya"}`)
	request := httptest.NewRequest("POST", "/login", body)
	request.Header.Set("Content-Type", "application/json")

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Login Success Nathan", string(bytes))

	/*
		=== RUN   TestRequestBody
		--- PASS: TestRequestBody (0.00s)
		PASS
	*/
}

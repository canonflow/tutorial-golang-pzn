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

type RegisterRequest struct {
	Username string `json:"username" xml:"username" form:"username"`
	Password string `json:"password" xml:"password" form:"password"`
	Name     string `json:"name" xml:"name" form:"name"`
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

func TestBodyParser(t *testing.T) {
	app.Post("/register", func(c *fiber.Ctx) error {
		request := new(RegisterRequest)
		err := c.BodyParser(request) // Parse the body into Struct
		if err != nil {
			return err
		}

		return c.SendString("Register Success " + request.Username)
	})
}

func TestBodyParserJSON(t *testing.T) {
	TestBodyParser(t)

	body := strings.NewReader(`{"username": "Nathan", "password": "Garzya", "name": "Nathan Garzya"}`)
	request := httptest.NewRequest("POST", "/register", body)
	request.Header.Set("Content-Type", "application/json")

	response, err := app.Test(request)
	assert.Nil(t, err)
	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Register Success Nathan", string(bytes))
	/*
		=== RUN   TestBodyParserJSON
		--- PASS: TestBodyParserJSON (0.00s)
		PASS
	*/
}

func TestBodyParserForm(t *testing.T) {
	TestBodyParser(t)

	body := strings.NewReader(`username=Nathan&password=Garzya&nameNathan Garzya`)
	request := httptest.NewRequest("POST", "/register", body)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := app.Test(request)
	assert.Nil(t, err)
	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Register Success Nathan", string(bytes))
	/*
		=== RUN   TestBodyParserForm
		--- PASS: TestBodyParserForm (0.00s)
		PASS
	*/
}

func TestBodyParserXML(t *testing.T) {
	TestBodyParser(t)

	body := strings.NewReader(
		`<RegisterRequest>
			<username>Nathan</username>
			<password>Garzya</password>
			<name>Nathan Garzya</name>
		</RegisterRequest>`)
	request := httptest.NewRequest("POST", "/register", body)
	request.Header.Set("Content-Type", "application/xml")

	response, err := app.Test(request)
	assert.Nil(t, err)
	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Register Success Nathan", string(bytes))
	/*
		=== RUN   TestBodyParserXML
		--- PASS: TestBodyParserXML (0.00s)
		PASS
	*/
}

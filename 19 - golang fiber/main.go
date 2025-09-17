package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		IdleTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
		ReadTimeout:  time.Second * 5,
		Prefork:      true, // Enable prefork mode
	})

	//! PREFORK
	if fiber.IsChild() {
		fmt.Println("This is a child process")
	} else {
		fmt.Println("This is the parent process")
	}

	//! MIDDLEWARE
	app.Use(func(c *fiber.Ctx) error {
		fmt.Println("Request received at:", time.Now())
		err := c.Next()
		fmt.Println("Response sent at:", time.Now())
		return err
	})

	//! MIDDLEWARE ON SPECIFIC ROUTE PREFIX
	app.Use("/api", func(c *fiber.Ctx) error {
		fmt.Println("API Request received at:", time.Now())
		err := c.Next()
		fmt.Println("API Response sent at:", time.Now())
		return err
	})

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello World!")
	})

	app.Get("/api/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello from API!")
	})

	err := app.Listen("localhost:3000")
	if err != nil {
		panic(err)
	}
}

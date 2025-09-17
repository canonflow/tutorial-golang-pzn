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

	if fiber.IsChild() {
		fmt.Println("This is a child process")
	} else {
		fmt.Println("This is the parent process")
	}

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello World!")
	})

	err := app.Listen("localhost:3000")
	if err != nil {
		panic(err)
	}
}

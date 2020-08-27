package main

import (
	"github.com/gofiber/fiber"
	"github.com/vasuvanka/gologger"
)

func main() {
	app := fiber.New()

	app.Use(gologger.Log)

	app.Get("/", func(c *fiber.Ctx) {
		c.Send("Hello, World 👋!")
	})

	app.Listen(3000)
}

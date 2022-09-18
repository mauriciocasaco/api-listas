package main

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"todolist/Listas"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		//return c.SendString("Hello, World!")
		c.Status(http.StatusOK).JSON(&fiber.Map{
			"message": "Estamos constrruyendo la funcion",
		})
		return nil
	})
	app.Post("/test", Listas.CreateList)

	app.Listen(":3000")
}

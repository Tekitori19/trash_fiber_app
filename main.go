package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./views")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("content", nil, "layouts/main")
	})

	app.Get("/sidebar", func(c *fiber.Ctx) error {
		return c.Render("partials/sidebar", nil)
	})

	app.Get("/main", func(c *fiber.Ctx) error {
		return c.Render("partials/main", nil)
	})

	app.Listen(":3000")
}

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
	app.Static("/", "./views/js")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("contents/CheckSNT/content", nil, "layouts/main")
	})

	app.Get("/UCLN_BCNN", func(c *fiber.Ctx) error {
		return c.Render("contents/UCLN_BCNN/content", nil, "layouts/main")
	})

	app.Get("/karaoke", func(c *fiber.Ctx) error {
		return c.Render("contents/karaoke/content", nil, "layouts/main")
	})

	app.Get("/tinhtiendien", func(c *fiber.Ctx) error {
		return c.Render("contents/tinhtiendien/content", nil, "layouts/main")
	})

	app.Get("/doitien", func(c *fiber.Ctx) error {
		return c.Render("contents/doitien/content", nil, "layouts/main")
	})

	app.Get("/tinhlainganhang", func(c *fiber.Ctx) error {
		return c.Render("contents/CheckSNT/content", nil, "layouts/main")
	})

	app.Get("/vaytienmuaxe", func(c *fiber.Ctx) error {
		return c.Render("contents/CheckSNT/content", nil, "layouts/main")
	})

	app.Get("/sapxepthongtinsv", func(c *fiber.Ctx) error {
		return c.Render("contents/CheckSNT/content", nil, "layouts/main")
	})

	app.Get("/fpoly_lott", func(c *fiber.Ctx) error {
		return c.Render("contents/CheckSNT/content", nil, "layouts/main")
	})

	app.Get("/phanso", func(c *fiber.Ctx) error {
		return c.Render("contents/CheckSNT/content", nil, "layouts/main")
	})

	app.Get("/sidebar", func(c *fiber.Ctx) error {
		return c.Render("partials/sidebar", nil)
	})

	app.Listen(":3000")
}

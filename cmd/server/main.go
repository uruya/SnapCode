package main

import (
	"snapcode/internal/render"

	"github.com/gofiber/fiber/v2"
)

type req struct {
	Code  string `json:"code"`
	Theme string `json:"theme"` // "dark"|"light"
}

func main() {
	app := fiber.New(fiber.Config{BodyLimit: 1024 * 64}) // 64 KB limit

	app.Post("/generate", func(c *fiber.Ctx) error {
		r := new(req)
		if err := c.BodyParser(&r); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "invalid json"})
		}

		if r.Theme == "" {
			r.Theme = "dark"
		}

		img, err := render.MakePNG(r.Code, r.Theme)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		c.Set("Content-Type", "image/png")
		return c.Send(img)
	})

	app.Get("/generate", func(c *fiber.Ctx) error {
		code := c.Query("code")
		theme := c.Query("theme", "dark") // default dark

		if code == "" {
			return c.Status(400).JSON(fiber.Map{"error": "`code` query param is required"})
		}
		img, err := render.MakePNG(code, theme)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		c.Set("Content-Type", "image/png")
		// Specify that CDN can cache for 1 day
		c.Set("Cache-Control", "public, max-age=86400")
		return c.Send(img)
	})

	app.Listen(":8080")
}

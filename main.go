package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	api := app.Group("/api")

	// 测试路由
	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("App running aaa bbbb ffff")
	})
	log.Fatal(app.Listen(":5000"))
}

package routes

import "github.com/gofiber/fiber"
func Setup() {
    app := fiber.New()
    app.Get("/", func(c *fiber.Ctx) {
        c.Send("Hello, World!")
    })
    app.Listen(3000)
}
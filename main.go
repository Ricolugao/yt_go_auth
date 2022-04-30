package main

import (
	"yt_go_auth/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	routes.Setup(app)

	app.Listen(":8000")
}

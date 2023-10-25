package main

import (
	"FluxStore/configs"
	"FluxStore/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	configs.ConnectDB()

	app.Static("/static", "./public")

	routes.UserRoute(app)

	app.Listen(":6000")
}

package main

import (
	"SideProject/Go-React-Demo/database"
	"SideProject/Go-React-Demo/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()

	app := fiber.New()

	routes.Setup(app)
	app.Listen(":3000")
}

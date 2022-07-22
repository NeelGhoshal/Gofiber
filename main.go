package main

import (
	"fiber/database"
	"fiber/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to this API")
}

func setupRoutes(app *fiber.App) {
	//welcome endpoint
	app.Get("/api", welcome)
	//user endpoints
	app.Post("/api/users", routes.CreateUser)

	app.Get("/api/users", routes.GetUsers)

	app.Get("/api/users/:id", routes.GetUser)

	app.Put("/api/users/:id", routes.UpdateUser)

	app.Delete("/api/users/:id", routes.DeleteUser)
}
func main() {
	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}

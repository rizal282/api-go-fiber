package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/rizal282/go-fiber-api/database"
	"github.com/rizal282/go-fiber-api/routes"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to my awesome API using Fiber Go")
}

func setupRoute(app *fiber.App) {
	// welcome endpoint
	app.Get("/api", welcome)

	// create user route
	app.Post("/api/users", routes.CreateUser)
	app.Get("/api/users", routes.GetUsers)
	app.Get("/api/users/:id", routes.GetUserById)
	app.Put("/api/users/:id", routes.UpdateUser)
	app.Delete("/api/users/:id", routes.DeleteUser)
}

func main() {
	// call connect db for init Database
	database.ConnectDb()

	// initial new App using fiber
	app := fiber.New()

	setupRoute(app)

	log.Fatal(app.Listen(":3000"))
}
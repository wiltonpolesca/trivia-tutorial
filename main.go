package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wiltonpolesca/go-trivia-tutorial/database"
)

func main () {
	database.ConnectDb()
	app := fiber.New()
	
	setupRoutes(app)
	app.Listen(":3000")
}
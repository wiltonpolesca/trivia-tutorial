package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wiltonpolesca/go-trivia-tutorial/handlers"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.ListFacts)
	app.Post("/fact", handlers.CreateFact)

	app.Get("/test", handlers.Home)

}

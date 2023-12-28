package router

import (
	"github.com/agustfricke/market-go/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
  app.Post("/signup", handlers.SignUp)
  app.Get("/", handlers.SignUpView)
}

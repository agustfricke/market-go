package main

import (
	"log"

	"github.com/agustfricke/market-go/database"
	"github.com/agustfricke/market-go/router"
	"github.com/gofiber/template/html/v2"

	"github.com/gofiber/fiber/v2"
)


func main() {
  database.ConnectDB()
  engine := html.New("./templates", ".html")
  app := fiber.New(fiber.Config{
    Views: engine, 
  })
  app.Static("/", "./public")
  router.SetupRoutes(app)
  log.Fatal(app.Listen(":42069"))
}

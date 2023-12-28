package handlers

import (
	"log"

	"github.com/agustfricke/market-go/database"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func SignInView(c *fiber.Ctx) error {
  return c.Render("sign_in", fiber.Map{})
}

func SignUpView(c *fiber.Ctx) error {
  return c.Render("sign_up", fiber.Map{})
}

func SignUp(c *fiber.Ctx) error {
  db := database.DB
	username := c.FormValue("username")
	password := c.FormValue("password")

	if username == "" || password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "El nombre de usuario y la contraseña son obligatorios",
		})
	}

	hashedPassword, hashErr := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if hashErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": hashErr.Error()})
	}

  id := uuid.New()

	_, err := db.Exec("INSERT INTO users(id, username, password) VALUES(?, ?, ?)", id, username, hashedPassword)
	if err != nil {
		log.Printf("Error al insertar usuario en la base de datos: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Ocurrió un error al registrar el usuario",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Usuario registrado exitosamente"})
}


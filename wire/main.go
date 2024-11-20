package main

import (
	"log"
	"net/http"
	"os"
	"wire/config"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env :", err.Error())
		os.Exit(1)
	}

	config.InitializeDB()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"message" : "oke",
		})
	})

	app.Listen(":3000")


	
}
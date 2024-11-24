package main

import (
	"log"
	"os"
	"wire/config"
	"wire/routes"

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

	routes.NewRoutesGroup(app).Install()

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000"
	}

	app.Listen(":" + port)

}

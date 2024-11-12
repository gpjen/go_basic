package main

import (
	"clean_arsitektur/internal/config"
	"clean_arsitektur/internal/delivery/http/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	config.InitDB()
	dbConn, err := config.DB.DB()
	if err != nil {
		panic("failed to connect to the database: " + err.Error())
	}
	defer dbConn.Close()

	app := fiber.New()

	routes.InitializeRoutes(app)

	app.Listen(":3000")
}

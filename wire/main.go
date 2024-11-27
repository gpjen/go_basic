package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
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

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000"
	}

	go func() {
		if err := app.Listen(":" + port); err != nil {
			fmt.Println("Error Starting Server :", err.Error())
		}
	}()

	<-stop
	fmt.Println("Shuting Dows Server...")

	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := app.Shutdown(); err != nil {
		fmt.Println("Fiber Shutdown Error :", err.Error())
	}

	dbInstance, _ := config.DB.DB()
	if err := dbInstance.Close(); err != nil {
		fmt.Println("Error To CLose Databse :", err.Error())
	}

	fmt.Println("Server gracefully stopped")

}

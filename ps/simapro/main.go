package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"github.com/gpjen/simapro/config"
	"github.com/gpjen/simapro/database"
	"github.com/gpjen/simapro/middlewares"
	"github.com/gpjen/simapro/routes"
	"github.com/gpjen/simapro/utils"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	config.LoadConfig()
	database.ConnectDB()
	utils.InitLogger()
	defer utils.LogFile.Close()

	app := fiber.New(fiber.Config{
		ReadBufferSize: 16384,
		BodyLimit:      50 * 1024 * 1024,
		AppName:        config.AppConfig.APP_NAME,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			r := utils.ResponseHandler{}
			return r.InternalServerError(c, []string{err.Error()})
		},
		CaseSensitive: true,
		ServerHeader:  "Fiber",
		ReadTimeout:   10 * time.Second,
		WriteTimeout:  10 * time.Second,
		IdleTimeout:   120 * time.Second,
		// StrictRouting: true,
		// Prefork:       true,
	})

	middlewares.SetupUserContext(app)
	middlewares.SetupLogger(app)
	middlewares.SetupCORS(app)
	routes.SetupRoutesApp(app)

	cronScheduler, err := utils.InitializeChronJobs()
	if err != nil {
		utils.Logger.Fatalf("Failed to initialize and start cron job: %v", err)
	}

	go func() {
		port := fmt.Sprintf(":%s", config.AppConfig.APP_PORT)
		log.Printf("🚀 Server running at http://localhost%s", port)
		if err := app.Listen(port); err != nil {
			log.Fatalf("❌ Failed to start server: %v", err)
		}
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
	utils.Logger.Println("Shutting down server...")
	cronScheduler.Stop()
	utils.Logger.Println("Cron job stopped...")

	if err := app.Shutdown(); err != nil {
		utils.Logger.Printf("Failed to shutdown server: %v", err)
	}
}

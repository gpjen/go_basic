package main

import (
	"clean_arsitektur/internal/config"
	"clean_arsitektur/internal/delivery/http/routes"
	"clean_arsitektur/internal/migrations"
	"flag"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Parse flags for migration
	migrationMode := flag.Bool("migrate", false, "Run database migrations")
	migrationAction := flag.String("action", "up", "Migration action: up, down, or status")
	flag.Parse()

	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Initialize database
	config.InitDB()
	dbConn, err := config.DB.DB()
	if err != nil {
		panic("failed to connect to the database: " + err.Error())
	}
	defer dbConn.Close()

	// Handle migration mode
	if *migrationMode {
		manager := migrations.NewMigrationManager(config.DB)

		// Initialize migration system
		if err := manager.Initialize(); err != nil {
			log.Printf("Error initializing migrations: %v\n", err)
			os.Exit(1)
		}

		// Execute requested migration action
		var migrationErr error
		switch *migrationAction {
		case "up":
			migrationErr = manager.Up()
		case "down":
			migrationErr = manager.Down()
		case "status":
			migrationErr = manager.Status()
		default:
			log.Printf("Unknown migration action: %s\n", *migrationAction)
			os.Exit(1)
		}

		if migrationErr != nil {
			log.Printf("Migration error: %v\n", migrationErr)
			os.Exit(1)
		}

		log.Println("Migration completed successfully")
		return
	}

	// Start web server if not in migration mode
	app := fiber.New(fiber.Config{
		// You can add your Fiber configurations here
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			// Custom error handling
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return c.Status(code).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	// Initialize routes
	routes.InitializeRoutes(app)

	// Get port from environment variable or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// Start server
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}

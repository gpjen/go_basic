package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file, make sure the file exists.")
		os.Exit(1)
	}

	app := fiber.New(fiber.Config{
		// DisableStartupMessage: true,
		BodyLimit: 100 * 1024 * 1024,
	})

	// Middleware untuk menangani CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Anda dapat menyesuaikan ini untuk membatasi domain yang diizinkan
	}))

	// limiter
	app.Use(limiter.New(limiter.Config{
		Max:        100,
		Expiration: time.Minute,
	}))

	// Middleware untuk melindungi dari serangan CSRF
	app.Use(csrf.New())

	// Middleware untuk meningkatkan keamanan dengan mengatur header HTTP
	app.Use(helmet.New())

	// Middleware untuk menggunakan ETag dalam respons
	app.Use(etag.New())

	// Middleware untuk mengompres respons
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // Atur level kompresi
	}))

	// Middleware untuk menyajikan file statis dari direktori "./assets"
	app.Use(filesystem.New(filesystem.Config{
		Root: http.Dir("./assets"),
	}))

	// Middleware untuk mencatat detail permintaan
	app.Use(logger.New(logger.Config{
		Format:     "${time} - ${ip} - ${method} ${path} - ${status} - ${latency} - ${user-agent} \n",
		TimeFormat: "2006-01-02 15:04:05",
	}))

	// midleware encrypt cookie
	appEncryptKey := os.Getenv("APP_ENCRYPT_KEY")
	if appEncryptKey == "" {
		fmt.Println("APP_ENCRYPT_KEY is not set. Exiting.")
		os.Exit(1)
	}
	encodedKey := base64.StdEncoding.EncodeToString([]byte(appEncryptKey))
	app.Use(encryptcookie.New(encryptcookie.Config{
		Key: encodedKey,
	}))

	// midleware request id
	app.Use(requestid.New())

	app.Get("/user", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"message": "Hello USER, World!",
		})
	})

	app.Get("/user/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"id": id,
		})
	})

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000"
	}
	err = app.Listen(":" + port)
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}

}

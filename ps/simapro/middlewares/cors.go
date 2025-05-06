package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gpjen/simapro/config"
)

func SetupCORS(app *fiber.App) {
	cfg := cors.Config{
		AllowOrigins: config.AppConfig.APP_ORIGINS,
		AllowMethods: "GET, POST, PUT, DELETE",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization, X-Identifier",
	}

	if cfg.AllowOrigins != "*" {
		cfg.AllowCredentials = true
	}

	app.Use(cors.New(cfg))
}

package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gpjen/go_basic/ps/mini_finance/services/api/client"
	"github.com/gpjen/go_basic/ps/mini_finance/services/api/handler"
)

func main() {

	app := fiber.New()

	client.InitGRPC()
	handler.RegisterRoutes(app)

	app.Listen(":3000")
}

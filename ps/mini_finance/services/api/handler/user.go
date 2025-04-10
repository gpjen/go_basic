package handler

import (
	"context"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gpjen/go_basic/ps/mini_finance/proto/user"
	"github.com/gpjen/go_basic/ps/mini_finance/services/api/client"
)

func RegisterRoutes(app *fiber.App) {
	app.Post("/users", CreateUser)
	app.Get("/users/:id", GetUser)
}

func CreateUser(c *fiber.Ctx) error {
	req := new(user.CreateUserRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).SendString("Invalid Input")
	}

	res, err := client.UserClient.CreateUser(context.Background(), req)
	if err != nil {
		return c.Status(500).SendString("Failed To Create User")
	}

	return c.JSON(res)
}

func GetUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	res, err := client.UserClient.GetUser(context.Background(), &user.GetUserRequest{
		Id: int32(id),
	})

	if err != nil {
		return c.Status(500).SendString("Failed to get user")
	}

	return c.JSON(res)

}

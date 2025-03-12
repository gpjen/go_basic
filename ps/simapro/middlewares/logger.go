package middlewares

import (
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gpjen/simapro/utils"
)

func SetupLogger(app *fiber.App) {

	app.Use(logger.New(logger.Config{
		Format: "[${user}] ${time}|${ip}|${status}|${method}|${path}| ${latency}\n",
		CustomTags: map[string]logger.LogFunc{
			"user": func(output logger.Buffer, c *fiber.Ctx, data *logger.Data, extraParam string) (int, error) {
				usrEmail := utils.GetUserSession(c).Email
				return output.WriteString(usrEmail)
			},
			"time": func(output logger.Buffer, c *fiber.Ctx, data *logger.Data, extraParam string) (int, error) {
				currentTime := time.Now().In(time.FixedZone("Asia/Jakarta", 7*3600))
				return output.WriteString(currentTime.Format("02-01-2006 15:04:05"))
			},
		},
		Output: utils.LogFile,
		Done: func(c *fiber.Ctx, logString []byte) {
			fmt.Print(string(logString))
		},
		Next: func(c *fiber.Ctx) bool {

			skippedPaths := []string{
				"/favicon.ico",
				"/swagger",
				// "/api/v1/file",
			}

			for _, path := range skippedPaths {
				if strings.HasPrefix(c.Path(), path) {
					return true
				}
			}

			return c.Method() == fiber.MethodOptions
		},
	}))

}

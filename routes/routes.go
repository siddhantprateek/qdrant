package routes

import (
	"qdrant/config"

	"github.com/gofiber/fiber/v2"
)

func apiRoutes(routes *fiber.App) {
	routes.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Qdrant App Server is Healthy.",
		})
	})
}

func Init() error {

	app := fiber.New()

	apiRoutes(app)

	port := config.GetEnviron("PORT")
	err := app.Listen(":" + port)
	return err
}

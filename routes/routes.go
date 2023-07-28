package routes

import (
	"qdrant/config"
	handlers "qdrant/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func apiRoutes(routes *fiber.App) {
	routes.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Qdrant App Server is Healthy.",
		})
	})

	routes.Get("/all", handlers.GetAllCollection)
}

func Init() error {

	app := fiber.New()
	app.Use(logger.New())
	apiRoutes(app)

	port := config.GetEnviron("PORT")
	err := app.Listen(":" + port)
	return err
}

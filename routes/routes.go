package routes

import (
	"qdrant/config"
	handlers "qdrant/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func ApiRoutes(routes *fiber.App) {
	routes.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Qdrant App Server is Healthy.",
		})
	})

	routes.Get("/all", handlers.GetAllCollection)
	routes.Post("/collection/create", handlers.CreateCollection)
	routes.Post("/field/create", handlers.CreateField)
	routes.Post("/upsert", handlers.AddVectorData)
	routes.Post("/data/id", handlers.RetrieveById)
	routes.Delete("/collection/delete", handlers.DeleteVectorCollection)
}

func Init() error {

	app := fiber.New()
	app.Use(logger.New())
	ApiRoutes(app)

	port := config.GetEnviron("PORT")
	err := app.Listen(":" + port)
	return err
}

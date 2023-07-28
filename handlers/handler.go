package handlers

import (
	db "qdrant/connection"

	"github.com/gofiber/fiber/v2"
	pb "github.com/qdrant/go-client/qdrant"
)

func GetAllCollection(c *fiber.Ctx) error {
	collectionsClient, ctx, cancel := db.QdrantDBConn()
	defer cancel()

	r, err := collectionsClient.List(ctx, &pb.ListCollectionsRequest{})
	if err != nil {
		return c.JSON(fiber.Map{
			"message": "Could not get collections",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Qdrant Collections",
		"data":    r.GetCollections(),
	})
}

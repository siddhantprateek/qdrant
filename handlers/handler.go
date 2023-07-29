package handlers

import (
	db "qdrant/connection"

	"github.com/gofiber/fiber/v2"
	pb "github.com/qdrant/go-client/qdrant"
)

func GetAllCollection(c *fiber.Ctx) error {
	_, collectionsClient, ctx, cancel := db.QdrantDBConn()
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

func CreateCollection(c *fiber.Ctx) error {
	type NewCollection struct {
		CollectionName string `json:"collectionName"`
	}
	newCollectionReq := new(NewCollection)

	if err := c.BodyParser(newCollectionReq); err != nil {
		return c.JSON(fiber.Map{
			"message": fiber.ErrBadRequest,
			"status":  fiber.StatusBadRequest,
		})
	}

	result, err := db.CreateQdCollection(newCollectionReq.CollectionName)
	if err != nil {
		return c.JSON(fiber.Map{
			"message": "Failed to create Collection: " + err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": result,
		"status":  fiber.StatusOK,
	})
}

func CreateField(c *fiber.Ctx) error {
	type CreateFieldReq struct {
		CollectionName string `json:"collectionName"`
		FieldName      string `json:"fieldName"`
	}

	newFieldBody := new(CreateFieldReq)
	if err := c.BodyParser(newFieldBody); err != nil {
		return c.JSON(fiber.Map{
			"message": fiber.ErrBadRequest,
			"status":  fiber.StatusBadRequest,
		})
	}

	result, err := db.CreateFieldIndex(newFieldBody.CollectionName, newFieldBody.FieldName)
	if err != nil {
		return c.JSON(fiber.Map{
			"message": "Failed to create Field: " + err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": result,
		"status":  fiber.StatusOK,
	})
}

// Insert vectors into a collection
func AddVectorData(c *fiber.Ctx) error {
	return nil
}

func DeleteVectorData(c *fiber.Ctx) error {
	return nil
}

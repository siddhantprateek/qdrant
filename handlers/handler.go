package handlers

import (
	"fmt"
	db "qdrant/connection"
	"qdrant/utils"

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
	type VectorPayload struct {
		Id             int    `json:"id"`
		City           string `json:"city"`
		Location       string `json:"location"`
		CollectionName string `json:"collectionName"`
	}

	newVectorPayload := new(VectorPayload)
	if err := c.BodyParser(newVectorPayload); err != nil {
		return c.JSON(fiber.Map{
			"message": fiber.ErrBadRequest,
			"status":  fiber.StatusBadRequest,
		})
	}

	waitUpsert := true
	upsertPoints := []*pb.PointStruct{
		{
			Id: &pb.PointId{
				PointIdOptions: &pb.PointId_Num{Num: uint64(newVectorPayload.Id)},
			},
			Vectors: &pb.Vectors{
				VectorsOptions: &pb.Vectors_Vector{Vector: &pb.Vector{Data: utils.RandomVector(4, 0.0, 1.0)}},
			},
			Payload: map[string]*pb.Value{
				"city": {
					Kind: &pb.Value_StringValue{StringValue: newVectorPayload.City},
				},
				"location": {
					Kind: &pb.Value_StringValue{StringValue: newVectorPayload.Location},
				},
			},
		},
	}
	conn, collections_client, ctx, cancel := db.QdrantDBConn()
	defer cancel()
	_ = collections_client

	pointsClient := pb.NewPointsClient(conn)
	_, err := pointsClient.Upsert(ctx, &pb.UpsertPoints{
		CollectionName: newVectorPayload.CollectionName,
		Wait:           &waitUpsert,
		Points:         upsertPoints,
	})
	if err != nil {
		return c.JSON(fiber.Map{
			"message": "Could not upsert points:" + err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"status": fiber.StatusCreated,
		"Upsert": len(upsertPoints),
	})
}

func RetrieveById(c *fiber.Ctx) error {
	type ByIdPayload struct {
		Id             int    `json:"id"`
		CollectionName string `json:"collectionName"`
	}

	newByIdPayload := new(ByIdPayload)
	if err := c.BodyParser(newByIdPayload); err != nil {
		return c.JSON(fiber.Map{
			"message": "[ERROR] Couldn't parse vector id:" + err.Error(),
		})
	}

	conn, collections_client, ctx, cancel := db.QdrantDBConn()
	defer cancel()
	_ = collections_client

	pointsClient := pb.NewPointsClient(conn)

	pointsById, err := pointsClient.Get(ctx, &pb.GetPoints{
		CollectionName: newByIdPayload.CollectionName,
		Ids: []*pb.PointId{
			{PointIdOptions: &pb.PointId_Num{
				Num: uint64(newByIdPayload.Id)}},
		},
	})
	if err != nil {
		return c.JSON(fiber.Map{
			"message": "Could not retrieve points:" + err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"status": fiber.StatusAccepted,
		"data":   pointsById.GetResult(),
	})
}

func DeleteVectorCollection(c *fiber.Ctx) error {
	type CollectionPayload struct {
		CollectionName string `json:"collectionName"`
	}
	newCollectionPayload := new(CollectionPayload)

	if err := c.BodyParser(newCollectionPayload); err != nil {
		return c.JSON(fiber.Map{
			"message": fiber.ErrBadRequest,
			"status":  fiber.StatusBadRequest,
		})
	}

	_, collections_client, ctx, cancel := db.QdrantDBConn()
	defer cancel()

	_, err := collections_client.Delete(ctx, &pb.DeleteCollection{
		CollectionName: newCollectionPayload.CollectionName,
	})
	if err != nil {
		return c.JSON(fiber.Map{
			"message": fmt.Sprintf("Could not delete collection: %s", newCollectionPayload.CollectionName),
		})
	}

	return c.JSON(fiber.Map{
		"status":  fiber.StatusAccepted,
		"message": fmt.Sprintf("Collection %s deleted.", newCollectionPayload.CollectionName),
	})
}

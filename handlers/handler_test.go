package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"qdrant/handlers"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetAllCollection(t *testing.T) {
	app := fiber.New()
	app.Get("/collections", handlers.GetAllCollection)

	req := httptest.NewRequest(http.MethodGet, "/collections", nil)
	resp, err := app.Test(req)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestCreateCollection(t *testing.T) {
	app := fiber.New()
	app.Post("/collections", handlers.CreateCollection)

	// request payload.
	payload := map[string]interface{}{
		"collectionName": "test_collection",
	}
	body, _ := json.Marshal(payload)
	req := httptest.NewRequest(http.MethodPost, "/collections", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

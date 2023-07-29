package routes_test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestRootRoute(t *testing.T) {
	app := fiber.New()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Test the response body
	expectedResponse := `{"message": "Qdrant App Server is Healthy.", "status": 200}`
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	assert.Equal(t, expectedResponse, string(body))
}

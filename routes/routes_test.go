package routes_test

import (
	"net/http/httptest"
	handlers "qdrant/handlers"
	routes "qdrant/routes"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestAPIRoutes(t *testing.T) {
	app := fiber.New()
	routes.ApiRoutes(app)
	expectedRoutes := []struct {
		method  string
		path    string
		handler fiber.Handler
	}{
		{method: "GET", path: "/", handler: nil},
		{method: "GET", path: "/all", handler: handlers.GetAllCollection},
		{method: "POST", path: "/collection/create", handler: handlers.CreateCollection},
		{method: "POST", path: "/field/create", handler: handlers.CreateField},
		{method: "POST", path: "/upsert", handler: handlers.AddVectorData},
		{method: "POST", path: "/data/id", handler: handlers.RetrieveById},
		{method: "DELETE", path: "/collection/delete", handler: handlers.DeleteVectorCollection},
	}

	for _, route := range expectedRoutes {
		route := route
		t.Run(route.method+" "+route.path, func(t *testing.T) {
			t.Parallel()
			req := httptest.NewRequest(route.method, route.path, nil)
			resp, err := app.Test(req)
			assert.Nil(t, err)
			assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		})
	}
}

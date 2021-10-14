package routes

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestUsersRoute(t *testing.T) {
	type args struct {
		route fiber.Router
	}
	tests := []struct {
		description  string // description of the test case
		route        string // route path to test
		expectedCode int    // expected HTTP status code
	}{

		// First test case
		{
			description:  "get HTTP status 200",
			route:        "/api/user",
			expectedCode: 200,
		},
		// Second test case
		{
			description:  "get HTTP status 404, when route is not exists",
			route:        "/not-found",
			expectedCode: 404,
		},
	}
	app := fiber.New()
	api := app.Group("/api")
	// Create route with GET method for test
	UsersRoute(api.Group("/user"))
	for _, tt := range tests {
		// Create a new http request with the route from the test case
		req := httptest.NewRequest("GET", tt.route, nil)

		// Perform the request plain with the app,
		// the second argument is a request latency
		// (set to -1 for no latency)
		resp, _ := app.Test(req, 5)

		// Verify, if the status code is as expected
		assert.Equalf(t, tt.expectedCode, resp.StatusCode, tt.description)
	}
}

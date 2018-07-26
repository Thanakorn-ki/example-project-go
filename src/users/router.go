package users

import (
	"net/http"

	"github.com/labstack/echo"
)

// Handlers
func createUser(c echo.Context) error {
	return nil
}

func findUser(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func updateUser(c echo.Context) error {
	return nil
}

func deleteUser(c echo.Context) error {
	return nil
}

func RouterUsers(api *echo.Group) {
	// Routes
	api.POST("/users", createUser)
	api.GET("/users", findUser)
	api.PUT("/users", updateUser)
	api.DELETE("/users", deleteUser)
}

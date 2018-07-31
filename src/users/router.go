package users

import (
	"net/http"

	"github.com/example-project-go/src/account"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

// Handlers
func createUser(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "createUser")
	}
}

func findUser(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var accounts []account.Account
		db.Table("account").Limit(3).Find(&accounts)
		return c.JSON(http.StatusOK, accounts)
	}
}

func updateUser(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "updateUser")
	}
}

func deleteUser(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "deleteUser")
	}
}

// RouterUsers is Init router user
func RouterUsers(api *echo.Group, db *gorm.DB) {
	// Routes
	api.POST("/users", createUser(db))
	api.GET("/users", findUser(db))
	api.PUT("/users", updateUser(db))
	api.DELETE("/users", deleteUser(db))
}

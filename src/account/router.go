package account

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

var DB *gorm.DB

func RouterAccount(api *echo.Group, db *gorm.DB) {
	DB = db
	api.POST("/account", CreateAccount)
	api.GET("/account", FetchAccounts)
}

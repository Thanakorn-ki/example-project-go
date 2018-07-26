package account

import (
	"net/http"

	"github.com/labstack/echo"
)

func CreateAccount(c echo.Context) (err error) {
	account := new(account)
	if err = c.Bind(account); err != nil {
		return
	}
	DB.Table("account").Create(&account)
	return c.JSON(http.StatusOK, account)
}

func FetchAccounts(c echo.Context) error {
	var accounts []account
	DB.Table("account").Find(&accounts)
	return c.JSON(http.StatusOK, accounts)
}

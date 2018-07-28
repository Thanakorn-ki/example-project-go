package main

import (
	"agp-example/src/account"
	"agp-example/src/config"
	"agp-example/src/users"

	"github.com/labstack/echo"
)

func init() {

}
func main() {
	db := config.Init()
	e := echo.New()
	e.Use(ServerHeader)
	api := e.Group("/api") // prefix /api
	users.RouterUsers(api, db)
	account.RouterAccount(api, db)

	e.Logger.Fatal(e.Start(":8005"))
}

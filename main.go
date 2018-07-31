package main

import (
	"github.com/example-project-go/src/account"
	"github.com/example-project-go/src/config"

	"github.com/example-project-go/src/users"

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

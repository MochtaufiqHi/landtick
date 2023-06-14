package main

import (
	"landtick/database"
	"landtick/pkg/mysql"
	"landtick/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	mysql.DatabaseInit()
	database.RunMigration()

	routes.RouteInit(e.Group("/api/v1"))

	e.Logger.Fatal(e.Start(":1323"))
}

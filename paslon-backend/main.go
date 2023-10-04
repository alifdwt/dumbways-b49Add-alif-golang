package main

import (
	"fmt"
	"myapp/database"
	"myapp/pkg/postgres"
	"myapp/routes"

	"github.com/labstack/echo/v4"
)



func main() {
	e := echo.New()

	postgres.DatabaseInit()
	database.RunMigration()

	routes.RouteInit(e.Group("/api/v1"))

	fmt.Println("Server running at localhost:8000")
	e.Logger.Fatal(e.Start("localhost:8000"))
}


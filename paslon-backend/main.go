package main

import (
	"fmt"
	"myapp/database"
	"myapp/pkg/postgres"
	"myapp/routes"

	"github.com/labstack/echo/v4"
	"github.com/rs/cors"
)



func main() {
	e := echo.New()

	postgres.DatabaseInit()
	database.RunMigration()
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"OPTIONS", "GET", "POST", "PUT"},
		AllowedHeaders: []string{"Content-Type", "X-CSRF-Token", "Authorization"},
		Debug:          true,
	})
	e.Use(echo.WrapMiddleware(corsMiddleware.Handler))

	routes.RouteInit(e.Group("/api/v1"))

	fmt.Println("Server running at localhost:8000")
	e.Logger.Fatal(e.Start("localhost:8000"))
}


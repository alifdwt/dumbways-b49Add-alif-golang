package routes

import (
	"myapp/handlers"
	"myapp/pkg/postgres"
	"myapp/repositories"

	"github.com/labstack/echo/v4"
	"github.com/rs/cors"
)

func PaslonRoutes(e *echo.Group) {
	paslonRepository := repositories.RepositoryPaslon(postgres.DB)
	h := handlers.HandlerPaslon(paslonRepository)

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"https://alifdwt.com", "https://www.google.com", "http://localhost:5173"},
		AllowedMethods: []string{"OPTIONS", "GET", "POST", "PUT"},
		AllowedHeaders: []string{"Content-Type", "X-CSRF-Token"},
		Debug:          true,
	})
	e.Use(echo.WrapMiddleware(corsMiddleware.Handler))

	e.GET("/paslons", h.FindPaslons)
	e.GET("/paslon/:paslonId", h.GetPaslon)
	e.POST("/paslon", h.CreatePaslon)
	e.PATCH("/paslon/:paslonId", h.UpdatePaslon)
	e.DELETE("/paslon/:paslonId", h.DeletePaslon)
}
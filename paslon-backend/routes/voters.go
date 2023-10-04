package routes

import (
	"myapp/handlers"
	"myapp/pkg/postgres"
	"myapp/repositories"

	"github.com/labstack/echo/v4"
	"github.com/rs/cors"
)

func VoterRoutes(e *echo.Group) {
	voterRepository := repositories.RepositoryVoter(postgres.DB)
	h := handlers.HandlerVoter(voterRepository)

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"https://alifdwt.com", "https://www.google.com", "http://localhost:5173"},
		AllowedMethods: []string{"OPTIONS", "GET", "POST", "PUT"},
		AllowedHeaders: []string{"Content-Type", "X-CSRF-Token"},
		Debug:          true,
	})
	e.Use(echo.WrapMiddleware(corsMiddleware.Handler))

	e.GET("/voters", h.FindVoter)
	e.GET("/voter/:voterId", h.GetVoter)
	e.POST("/voter", h.CreateVoter)
	e.DELETE("/voter/:voterId", h.DeleteVoter)
}
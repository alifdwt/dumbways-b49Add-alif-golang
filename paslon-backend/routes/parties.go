package routes

import (
	"myapp/handlers"
	"myapp/pkg/postgres"
	"myapp/repositories"

	"github.com/labstack/echo/v4"
	"github.com/rs/cors"
)

func PartyRoutes(e *echo.Group) {
	partyRepository := repositories.RepositoryParty(postgres.DB)
	h := handlers.HandlerParty(partyRepository)

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"OPTIONS", "GET", "POST", "PUT"},
		AllowedHeaders: []string{"Content-Type", "X-CSRF-Token", "Authorization"},
		Debug:          true,
	})
	e.Use(echo.WrapMiddleware(corsMiddleware.Handler))

	e.GET("/parties", h.FindParties)
	e.GET("/party/:partyId", h.GetParty)
	e.POST("/party", h.CreateParty)
	e.PATCH("/party/:partyId", h.UpdateParty)
	e.DELETE("/party/:partyId", h.DeleteParty)
}
package routes

import (
	"myapp/handlers"
	"myapp/pkg/postgres"
	"myapp/repositories"

	"github.com/labstack/echo/v4"
)

func PartyRoutes(e *echo.Group) {
	partyRepository := repositories.RepositoryParty(postgres.DB)
	h := handlers.HandlerParty(partyRepository)

	e.GET("/parties", h.FindParties)
	e.GET("/party/:partyId", h.GetParty)
	e.POST("/party", h.CreateParty)
	e.PATCH("/party/:partyId", h.UpdateParty)
	e.DELETE("/party/:partyId", h.DeleteParty)
}
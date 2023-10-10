package routes

import (
	"myapp/handlers"
	"myapp/pkg/postgres"
	"myapp/repositories"

	"github.com/labstack/echo/v4"
)

func VoterRoutes(e *echo.Group) {
	voterRepository := repositories.RepositoryVoter(postgres.DB)
	h := handlers.HandlerVoter(voterRepository)

	e.GET("/voters", h.FindVoter)
	e.GET("/voter/:voterId", h.GetVoter)
	e.POST("/voter", h.CreateVoter)
	e.DELETE("/voter/:voterId", h.DeleteVoter)
}
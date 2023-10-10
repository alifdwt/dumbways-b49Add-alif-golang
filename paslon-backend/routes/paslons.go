package routes

import (
	"myapp/handlers"
	"myapp/pkg/postgres"
	"myapp/repositories"

	"github.com/labstack/echo/v4"
)

func PaslonRoutes(e *echo.Group) {
	paslonRepository := repositories.RepositoryPaslon(postgres.DB)
	h := handlers.HandlerPaslon(paslonRepository)

	e.GET("/paslons", h.FindPaslons)
	e.GET("/paslon/:paslonId", h.GetPaslon)
	e.POST("/paslon", h.CreatePaslon)
	e.PATCH("/paslon/:paslonId", h.UpdatePaslon)
	e.DELETE("/paslon/:paslonId", h.DeletePaslon)
}
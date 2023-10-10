package routes

import (
	"myapp/handlers"
	"myapp/pkg/postgres"
	"myapp/repositories"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Group) {
	userRepository := repositories.RepositoryUser(postgres.DB)
	h := handlers.HandlerUser(userRepository)

	e.POST("/auth/register", h.Register)
	e.POST("/auth/login", h.Login)
	e.GET("/auth/check", h.Check)
}
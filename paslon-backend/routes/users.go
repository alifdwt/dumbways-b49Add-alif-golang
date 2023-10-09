package routes

import (
	"myapp/handlers"
	"myapp/pkg/postgres"
	"myapp/repositories"

	"github.com/labstack/echo/v4"
	"github.com/rs/cors"
)

func UserRoutes(e *echo.Group) {
	userRepository := repositories.RepositoryUser(postgres.DB)
	h := handlers.HandlerUser(userRepository)

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"OPTIONS", "GET", "POST", "PUT"},
		AllowedHeaders: []string{"Content-Type", "X-CSRF-Token", "Authorization"},
		Debug:          true,
	})
	e.Use(echo.WrapMiddleware(corsMiddleware.Handler))

	e.POST("/auth/register", h.Register)
	e.POST("/auth/login", h.Login)
	e.GET("/auth/check", h.Check)
}
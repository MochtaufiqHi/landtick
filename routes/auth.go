package routes

import (
	handlers "landtick/handlers"
	"landtick/pkg/mysql"
	"landtick/repository"

	"github.com/labstack/echo/v4"
)

func AuthRoutes(e *echo.Group) {
	authRepository := repository.RepositoryAuth(mysql.DB)
	h := handlers.HandlersAuth(authRepository)

	e.POST("/register", h.Register)
	e.POST("/login", h.Login)
}

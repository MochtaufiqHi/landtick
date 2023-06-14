package routes

import (
	handlers "landtick/handlers"
	"landtick/pkg/mysql"
	"landtick/repository"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Group) {
	userRepository := repository.RepositoryUser(mysql.DB)
	h := handlers.HandlersUser(userRepository)

	e.GET("/user", h.FindUser)
	e.GET("/user/:id", h.GetUser)
}

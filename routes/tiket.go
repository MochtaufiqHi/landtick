package routes

import (
	"landtick/handlers"
	"landtick/pkg/middleware"
	"landtick/pkg/mysql"
	"landtick/repository"

	"github.com/labstack/echo/v4"
)

func TiketRoutes(e *echo.Group) {
	tiketRepository := repository.RepositoryTiket(mysql.DB)
	h := handlers.HandlersTiket(tiketRepository)

	e.POST("/tiket", middleware.Auth(h.CreateTiket))
	e.GET("/tiket", h.FindTiket)
	e.GET("/tiket/:id", h.GetTiket)
}

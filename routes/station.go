package routes

import (
	handlers "landtick/handlers"
	"landtick/pkg/mysql"
	"landtick/repository"

	"github.com/labstack/echo/v4"
)

func StationRoutes(e *echo.Group) {
	stationRepository := repository.RepositoryStation(mysql.DB)
	h := handlers.HandlersStation(stationRepository)

	e.POST("/station", h.CreateStation)
	e.GET("/station", h.FindStation)
	e.GET("/station/:id", h.GetStation)
	e.GET("/station/:name", h.GetStasionByName)
}

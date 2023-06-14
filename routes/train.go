package routes

import (
	handlers "landtick/handlers"
	"landtick/pkg/mysql"
	"landtick/repository"

	"github.com/labstack/echo/v4"
)

func TrainRoutes(e *echo.Group) {
	trainRepository := repository.RepositoryTrain(mysql.DB)
	h := handlers.HandlersTrain(trainRepository)

	e.POST("/train", h.CreateTrain)
	e.GET("/train", h.FindTrain)
	e.GET("/kereta/:id", h.GetTrain)
}

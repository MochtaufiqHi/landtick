package routes

import (
	"landtick/handlers"
	"landtick/pkg/mysql"
	"landtick/repository"

	"github.com/labstack/echo/v4"
)

func TransaksiRoutes(e *echo.Group) {
	transaksiRepository := repository.RepositoryTransaksi(mysql.DB)
	h := handlers.HandlersTransaksi(transaksiRepository)

	e.POST("/transaksi", h.CreateTransaksi)
	e.GET("/transaksi", h.FindTransaksi)
	e.GET("/transaksi/:id", h.GetTransaksi)
}

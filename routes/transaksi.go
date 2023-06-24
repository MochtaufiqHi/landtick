package routes

import (
	"landtick/handlers"
	"landtick/pkg/middleware"
	"landtick/pkg/mysql"
	"landtick/repository"

	"github.com/labstack/echo/v4"
)

func TransaksiRoutes(e *echo.Group) {
	transaksiRepository := repository.RepositoryTransaksi(mysql.DB)
	h := handlers.HandlersTransaksi(transaksiRepository)

	e.POST("/transaksi", middleware.Auth(h.CreateTransaksi))
	e.GET("/transaksi", h.FindTransaksi)
	e.GET("/transaksi-user/:id", h.FindTransaksiByUserId)
	e.GET("/transaksi/:id", h.GetTransaksi)
	e.POST("/notification", h.Notification)
}

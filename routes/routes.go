package routes

import "github.com/labstack/echo/v4"

func RouteInit(e *echo.Group) {
	AuthRoutes(e)
	TrainRoutes(e)
	// KotaRoutes(e)
	// StasiunRoutes(e)
	// StasiunTujuanRoutes(e)
	TiketRoutes(e)
	TransaksiRoutes(e)
	UserRoutes(e)
	// KeretaRoutes(e)
}

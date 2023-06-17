package handlers

import (
	dto "landtick/dto/result"
	stationdto "landtick/dto/station"
	"landtick/models"
	"landtick/repository"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type stationHandlers struct {
	StationRepository repository.StationRepository
}

func HandlersStation(StationRepository repository.StationRepository) *stationHandlers {
	return &stationHandlers{StationRepository}
}

func (h *stationHandlers) FindStation(c echo.Context) error {
	station, err := h.StationRepository.FindStation()

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: station})
}

func (h *stationHandlers) GetStation(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	station, err := h.StationRepository.GetStation(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: station})
}

func (h *stationHandlers) CreateStation(c echo.Context) error {
	request := new(stationdto.StationRequest)

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	station := models.Station{
		Name: request.Name,
		Code: request.Code,
		Kota: request.Kota,
	}
	data, err := h.StationRepository.CreateStation(station)
	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: data})

}

func (h *stationHandlers) GetStasionByName(c echo.Context) error {
	name := c.Param("name")

	station, err := h.StationRepository.GetStasionByName(name)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: station})
}

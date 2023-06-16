package handlers

import (
	dto "landtick/dto/result"
	tiketdto "landtick/dto/tiket"
	"landtick/models"
	"landtick/repository"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type tiketHandlers struct {
	TiketRepository repository.TiketRepository
}

func HandlersTiket(TiketRepository repository.TiketRepository) *tiketHandlers {
	return &tiketHandlers{TiketRepository}
}

func (h *tiketHandlers) FindTiket(c echo.Context) error {
	tiket, err := h.TiketRepository.FindTiket()

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: tiket})
}

func (h *tiketHandlers) GetTiket(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	tiket, err := h.TiketRepository.GetTiket(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: tiket})
}

func (h *tiketHandlers) CreateTiket(c echo.Context) error {
	request := new(tiketdto.TiketRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: "Validator error"})
	}

	tiket := models.Tiket{
		Name:         request.Name,
		TrainID:      request.TrainID,
		Train:        models.TrainResponse{},
		JamBerangkat: request.JamBerangkat,
		JamTiba:      request.JamTiba,
		Durasi:       request.Durasi,
		Harga:        request.Harga,
		Tanggal:      request.Tanggal,
		Kuota:        request.Kuota,
		StasiunAwal:  request.StasiunAwal,
		StasiunAkhir: request.StasiunAkhir,
	}

	data, err := h.TiketRepository.CreateTiket(tiket)
	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponseTiket(data)})
}

func convertResponseTiket(u models.Tiket) tiketdto.TiketRespon {
	return tiketdto.TiketRespon{
		ID:      u.ID,
		Name:    u.Name,
		TrainID: u.TrainID,
		// JenisKereta:  models.KeretaRespon(u.JenisKereta),
		JamBerangkat: u.JamBerangkat,
		JamTiba:      u.JamTiba,
		Durasi:       u.Durasi,
		Harga:        u.Harga,
		Tanggal:      u.Tanggal,
		Kuota:        u.Kuota,
		StasiunAwal:  u.StasiunAwal,
		StasiunAkhir: u.StasiunAkhir,
	}
}

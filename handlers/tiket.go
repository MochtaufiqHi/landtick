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
	// request := new(tiketdto.TiketRequest)
	// if err := c.Bind(request); err != nil {
	// 	return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	// }

	train_id, _ := strconv.Atoi(c.FormValue("train_id"))
	harga, _ := strconv.Atoi(c.FormValue("harga"))
	kuota, _ := strconv.Atoi(c.FormValue("kuota"))

	request := tiketdto.TiketRequest{
		Name:         c.FormValue("name"),
		JamBerangkat: c.FormValue("jam_berangkat"),
		JamTiba:      c.FormValue("jam_tiba"),
		StasiunAwal:  c.FormValue("stasiun_awal"),
		StasiunAkhir: c.FormValue("stasiun_akhir"),
		Durasi:       c.FormValue("durasi"),
		Tanggal:      c.FormValue("tanggal"),
		TrainID:      train_id,
		Harga:        harga,
		Kuota:        kuota,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: "Validator error"})
	}

	// userLogin := c.Get("userLogin")
	// userId := userLogin.(jwt.MapClaims)["id"].(float64)

	idTrain, _ := h.TiketRepository.GetTiketByID(request.TrainID)

	tiket := models.Tiket{
		Name:         request.Name,
		TrainID:      request.TrainID,
		Train:        idTrain,
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
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}
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

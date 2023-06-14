package handlers

import (
	dto "landtick/dto/result"
	transaksidto "landtick/dto/transaksi"
	"landtick/models"
	"landtick/repository"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type transaksiHandlers struct {
	TransaksiRepository repository.TransaksiRepository
}

func HandlersTransaksi(TransaksiRepository repository.TransaksiRepository) *transaksiHandlers {
	return &transaksiHandlers{TransaksiRepository}
}

func (h *transaksiHandlers) FindTransaksi(c echo.Context) error {
	transaksi, err := h.TransaksiRepository.FindTransaksi()

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: transaksi})
}

func (h *transaksiHandlers) GetTransaksi(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	transaksi, err := h.TransaksiRepository.GetTransaksi(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: transaksi})
}

func (h *transaksiHandlers) CreateTransaksi(c echo.Context) error {
	request := new(transaksidto.TransaksiRequest)
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	transaksi := models.Transaksi{
		Qty:        request.Qty,
		Total:      request.Total,
		Status:     request.Status,
		Attachment: request.Attachment,
		UserID:     request.UserID,
		User:       models.UserRespon{},
		TiketID:    request.TiketID,
		Tiket:      models.TiketRespon{},
	}

	data, err := h.TransaksiRepository.CreateTransaksi(transaksi)
	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponseTransaksi(data)})
}

func convertResponseTransaksi(u models.Transaksi) transaksidto.TransaksiRespon {
	return transaksidto.TransaksiRespon{
		ID:      u.ID,
		Qty:     u.Qty,
		Total:   u.Total,
		Status:  u.Status,
		UserID:  u.UserID,
		TiketID: u.TiketID,
	}
}

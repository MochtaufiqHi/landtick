package handlers

import (
	dto "landtick/dto/result"
	traindto "landtick/dto/train"
	"landtick/models"
	"landtick/repository"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type trainHandlers struct {
	TrainRepository repository.TrainRepository
}

func HandlersTrain(TrainRepository repository.TrainRepository) *trainHandlers {
	return &trainHandlers{TrainRepository}
}

func (h *trainHandlers) FindTrain(c echo.Context) error {
	train, err := h.TrainRepository.FindTrain()

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: train})

}

func (h *trainHandlers) GetTrain(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	Train, err := h.TrainRepository.GetTrain(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: Train})
}

func (h *trainHandlers) CreateTrain(c echo.Context) error {
	request := new(traindto.TrainRequest)

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	train := models.Train{
		Name: request.Name,
	}

	data, err := h.TrainRepository.CreateTrain(train)
	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: data})
}

package handlers

import (
	dto "landtick/dto/result"
	"landtick/repository"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type userHandlers struct {
	UserRepository repository.UserRepository
}

func HandlersUser(UserRepository repository.UserRepository) *userHandlers {
	return &userHandlers{UserRepository}
}

func (h *userHandlers) FindUser(c echo.Context) error {
	user, err := h.UserRepository.FindUser()

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: user})

}

func (h *userHandlers) GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	User, err := h.UserRepository.GetUser(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: User})
}

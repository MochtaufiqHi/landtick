package handlers

import (
	authdto "landtick/dto/auth"
	dto "landtick/dto/result"
	"landtick/models"
	"landtick/pkg/bcrypt"
	jwtToken "landtick/pkg/jwt"
	"landtick/repository"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	// "golang.org/x/crypto/bcrypt"
)

type handlersAuth struct {
	AuthRepository repository.AuthRepository
}

func HandlersAuth(AuthRepository repository.AuthRepository) *handlersAuth {
	return &handlersAuth{AuthRepository}
}

func (h *handlersAuth) Register(c echo.Context) error {
	// membuat alokasi memori untuk auth request
	request := new(authdto.AuthRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: "Binding Gagal"})
	}

	// hashing password
	password, err := bcrypt.HashingPassword(request.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: "Hasshing password gagal"})
	}

	// validasi bidang struct
	validation := validator.New()
	err = validation.Struct(request)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: "Validator Gagal"})
	}

	user := models.User{
		Fullname: request.Fullname,
		Username: request.Username,
		Email:    request.Email,
		Password: password,
		Gender:   request.Gender,
		Phone:    request.Phone,
		Address:  request.Address,
		Role:     "user",
	}

	data, err := h.AuthRepository.Register(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: "Create Registrasi gagal"})
	}

	dataRespon := authdto.AuthRespon{
		Fullname: data.Fullname,
		Username: data.Username,
		Email:    data.Email,
		Password: data.Password,
		Gender:   data.Gender,
		Phone:    data.Phone,
		Address:  data.Address,
		Role:     data.Role,
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: dataRespon})
}

func (h *handlersAuth) Login(c echo.Context) error {
	// membuat alokasi memori dto auth request
	request := new(authdto.AuthRequest)
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	user := models.User{
		Email:    request.Email,
		Password: request.Password,
	}

	// check email
	dataUser, err := h.AuthRepository.Login(user.Email)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	validasiUser := bcrypt.CheckPassword(request.Password, dataUser.Password)
	if !validasiUser {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	// generate token
	claims := jwt.MapClaims{}
	claims["id"] = user.ID
	claims["role"] = user.Role
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()

	token, errGenerateToken := jwtToken.GenerateToken(&claims)

	// token, errGenerateToken := jwtToken.GenerateToken(&claims)
	if errGenerateToken != nil {
		log.Println(errGenerateToken)
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	loginRespon := authdto.LoginRespon{
		ID:       dataUser.ID,
		Fullname: dataUser.Fullname,
		Username: dataUser.Username,
		Email:    dataUser.Email,
		Password: dataUser.Password,
		Gender:   dataUser.Gender,
		Phone:    dataUser.Phone,
		Address:  dataUser.Address,
		Role:     dataUser.Role,
		Token:    token,
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: loginRespon})
}

func (h *handlersAuth) CheckAuth(c echo.Context) error {
	userLogin := c.Get("userLogin")
	userID := userLogin.(jwt.MapClaims)["id"].(float64)

	user, _ := h.AuthRepository.CheckAuth(int(userID))

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: user})
}

package handlers

import (
	"fmt"
	dto "landtick/dto/result"
	transaksidto "landtick/dto/transaksi"
	"landtick/models"
	"landtick/repository"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"gopkg.in/gomail.v2"
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

func (h *transaksiHandlers) FindTransaksiByUserId(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	transasksi, err := h.TransaksiRepository.GetTransaksiByUserId(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: transasksi})
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

	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)

	// userId := 3
	request.UserID = int(userId)
	request.Status = "pending"

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	var transactionIsMatch = false
	var transactionId int
	for !transactionIsMatch {
		transactionId = int(time.Now().Unix())
		transactionData, _ := h.TransaksiRepository.GetTransaksi(transactionId)
		if transactionData.ID == 0 {
			transactionIsMatch = true
		}
	}

	transaksi := models.Transaksi{
		ID:         transactionId,
		Qty:        request.Qty,
		Total:      request.Total,
		Status:     request.Status,
		Attachment: request.Attachment,
		UserID:     int(userId),
		User:       models.UserRespon{},
		TiketID:    request.TiketID,
		Tiket:      models.TiketRespon{},
	}

	dataTransactions, err := h.TransaksiRepository.CreateTransaksi(transaksi)
	// return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponseTransaksi(data)})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	// 1. Initiate Snap client
	var s = snap.Client{}
	s.New(os.Getenv("SERVER_KEY"), midtrans.Sandbox)
	// Use to midtrans.Production if you want Production Environment (accept real transaction).

	// 2. Initiate Snap request param
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(dataTransactions.ID),
			GrossAmt: int64(dataTransactions.Total),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: dataTransactions.User.Fullname,
			Email: dataTransactions.User.Email,
		},
	}

	// 3. Execute request create Snap transaction to Midtrans Snap API
	snapResp, _ := s.CreateTransaction(req)

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: snapResp})
}

func (h *transaksiHandlers) Notification(c echo.Context) error {
	var notificationPayload map[string]interface{}

	if err := c.Bind(&notificationPayload); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	transactionStatus := notificationPayload["transaction_status"].(string)
	fraudStatus := notificationPayload["fraud_status"].(string)
	orderId := notificationPayload["order_id"].(string)

	ID, _ := strconv.Atoi(orderId)

	fmt.Print("ini Payload nya", notificationPayload)

	transaction, _ := h.TransaksiRepository.GetTransaksi(ID)

	if transactionStatus == "capture" {
		if fraudStatus == "challenge" {
			// TODO set transaction status on your database to 'challenge'
			// e.g: 'Payment status challenged. Please take action on your Merchant Administration Portal
			h.TransaksiRepository.UpdateTransaksi("pending", ID)
		} else if fraudStatus == "accept" {
			// TODO set transaction status on your database to 'success'
			SendMail("success", transaction)
			h.TransaksiRepository.UpdateTransaksi("success", ID)
		}
	} else if transactionStatus == "settlement" {
		// TODO set transaction status on your databaase to 'success'
		h.TransaksiRepository.UpdateTransaksi("success", ID)
	} else if transactionStatus == "deny" {
		// TODO you can ignore 'deny', because most of the time it allows payment retries
		// and later can become success
		h.TransaksiRepository.UpdateTransaksi("failed", ID)
	} else if transactionStatus == "cancel" || transactionStatus == "expire" {
		// TODO set transaction status on your databaase to 'failure'
		h.TransaksiRepository.UpdateTransaksi("failed", ID)
	} else if transactionStatus == "pending" {
		// SendMail("success", transaction)
		// TODO set transaction status on your databaase to 'pending' / waiting payment
		h.TransaksiRepository.UpdateTransaksi("pending", ID)
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: notificationPayload})
}

func SendMail(status string, transaction models.Transaksi) {

	if status != transaction.Status && (status == "success") {
		var CONFIG_SMTP_HOST = "smtp.gmail.com"
		var CONFIG_SMTP_PORT = 587
		var CONFIG_SENDER_NAME = "DumbTour <demo.dumbways@gmail.com>"
		var CONFIG_AUTH_EMAIL = os.Getenv("EMAIL_SYSTEM")
		var CONFIG_AUTH_PASSWORD = os.Getenv("PASSWORD_SYSTEM")

		var productName = transaction.Tiket.Name
		var price = strconv.Itoa(transaction.Total)

		mailer := gomail.NewMessage()
		mailer.SetHeader("From", CONFIG_SENDER_NAME)
		mailer.SetHeader("To", transaction.User.Email)
		mailer.SetHeader("Subject", "Transaction Status")
		mailer.SetBody("text/html", fmt.Sprintf(`<!DOCTYPE html>
    <html lang="en">
      <head>
      <meta charset="UTF-8" />
      <meta http-equiv="X-UA-Compatible" content="IE=edge" />
      <meta name="viewport" content="width=device-width, initial-scale=1.0" />
      <title>Document</title>
      <style>
        h1 {
        color: brown;
        }
      </style>
      </head>
      <body>
      <h2>Product payment :</h2>
      <ul style="list-style-type:none;">
        <li>Name : %s</li>
        <li>Total payment: Rp.%s</li>
        <li>Status : <b>%s</b></li>
      </ul>
      </body>
    </html>`, productName, price, status))

		dialer := gomail.NewDialer(
			CONFIG_SMTP_HOST,
			CONFIG_SMTP_PORT,
			CONFIG_AUTH_EMAIL,
			CONFIG_AUTH_PASSWORD,
		)

		err := dialer.DialAndSend(mailer)
		if err != nil {
			log.Fatal(err.Error())
		}

		log.Println("Mail sent! to " + transaction.User.Email)
	}
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

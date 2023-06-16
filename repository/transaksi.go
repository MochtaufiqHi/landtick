package repository

import (
	"landtick/models"

	"gorm.io/gorm"
)

type TransaksiRepository interface {
	CreateTransaksi(transaksi models.Transaksi) (models.Transaksi, error)
	FindTransaksi() ([]models.Transaksi, error)
	GetTransaksi(ID int) (models.Transaksi, error)
	UpdateTransaksi(status string, orderId int) (models.Transaksi, error)
}

func RepositoryTransaksi(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateTransaksi(transaksi models.Transaksi) (models.Transaksi, error) {
	err := r.db.Create(&transaksi).Error

	return transaksi, err
}

func (r *repository) FindTransaksi() ([]models.Transaksi, error) {
	var transaksi []models.Transaksi

	err := r.db.Preload("User").Preload("Tiket.Train").Find(&transaksi).Error

	return transaksi, err
}

func (r *repository) GetTransaksi(ID int) (models.Transaksi, error) {
	var transaksi models.Transaksi

	err := r.db.Preload("User").Preload("Tiket.Train").First(&transaksi, ID).Error

	return transaksi, err
}

func (r *repository) UpdateTransaksi(status string, orderId int) (models.Transaksi, error) {
	var transaksi models.Transaksi
	r.db.Preload("User").Preload("Tiket.Train").First(&transaksi, orderId)

	if status != transaksi.Status && status == "success" {
		var tiket models.Tiket
		r.db.First(&tiket, transaksi.Tiket.ID)
		tiket.Kuota = tiket.Kuota - transaksi.Qty
		r.db.Save(&tiket)
	}

	// fmt.Println("ini transaction counter qty", transaction.CounterQty)

	transaksi.Status = status
	err := r.db.Save(&transaksi).Error
	return transaksi, err
}

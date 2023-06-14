package repository

import (
	"landtick/models"

	"gorm.io/gorm"
)

type TransaksiRepository interface {
	CreateTransaksi(transaksi models.Transaksi) (models.Transaksi, error)
	FindTransaksi() ([]models.Transaksi, error)
	GetTransaksi(ID int) (models.Transaksi, error)
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

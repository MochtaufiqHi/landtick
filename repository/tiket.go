package repository

import (
	"landtick/models"

	"gorm.io/gorm"
)

type TiketRepository interface {
	CreateTiket(tiket models.Tiket) (models.Tiket, error)
	FindTiket() ([]models.Tiket, error)
	GetTiket(ID int) (models.Tiket, error)
}

func RepositoryTiket(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateTiket(tiket models.Tiket) (models.Tiket, error) {
	err := r.db.Create(&tiket).Error

	return tiket, err
}

func (r *repository) FindTiket() ([]models.Tiket, error) {
	var tiket []models.Tiket

	err := r.db.Preload("Train").Find(&tiket).Error

	return tiket, err
}

func (r *repository) GetTiket(ID int) (models.Tiket, error) {
	var tiket models.Tiket

	err := r.db.Preload("Train").First(&tiket, ID).Error

	return tiket, err
}

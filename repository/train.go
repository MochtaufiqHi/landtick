package repository

import (
	"landtick/models"

	"gorm.io/gorm"
)

type TrainRepository interface {
	CreateTrain(train models.Train) (models.Train, error)
	FindTrain() ([]models.Train, error)
	GetTrain(ID int) (models.Train, error)
}

func RepositoryTrain(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateTrain(train models.Train) (models.Train, error) {
	err := r.db.Create(&train).Error

	return train, err
}

func (r *repository) FindTrain() ([]models.Train, error) {
	var train []models.Train
	err := r.db.Find(&train).Error

	return train, err
}

func (r *repository) GetTrain(ID int) (models.Train, error) {
	var train models.Train
	err := r.db.First(&train, ID).Error

	return train, err
}

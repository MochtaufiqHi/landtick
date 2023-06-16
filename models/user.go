package models

import "time"

type User struct {
	ID          int                 `json:"id"`
	Fullname    string              `json:"fullname" gorm:"type: varchar(255)"`
	Username    string              `json:"username" gorm:"type: varchar(255)"`
	Email       string              `json:"email" gorm:"type: varchar(255)"`
	Password    string              `json:"password" gorm:"type: varchar(255)"`
	Gender      string              `json:"gender" gorm:"type: varchar(255)"`
	Phone       string              `json:"phone" gorm:"type: varchar(255)"`
	Address     string              `json:"address" gorm:"type: varchar(255)"`
	Role        string              `json:"role" gorm:"type: varchar(255)"`
	TransaksiID int                 `json:"transaksi_id"`
	Transaksi   []TransaksiResponse `json:"transaksi"`
	CreatedAt   time.Time           `json:"-"`
	UpdatedAt   time.Time           `json:"-"`
}

type UserRespon struct {
	ID          int    `json:"id"`
	Fullname    string `json:"fullname"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Gender      string `json:"gender"`
	Phone       string `json:"phone"`
	Address     string `json:"address"`
	Role        string `json:"role"`
	TransaksiID int    `json:"transaksi_id"`
	// Transaksi   []TransaksiResponse `json:"transaksi"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func (UserRespon) TableName() string {
	return "users"
}

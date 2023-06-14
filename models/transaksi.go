package models

import "time"

type Transaksi struct {
	ID         int         `json:"id"`
	Qty        int         `json:"qty" gorm:"type: int"`
	Total      int         `json:"total" gorm:"type: int"`
	Status     string      `json:"status" gorm:"type: varchar(255)"`
	Attachment string      `json:"attachment" gorm:"type: varchar(255)"`
	UserID     int         `json:"user_id" gorm:"type: int"`
	User       UserRespon  `json:"user" gorm:"foreignKey:UserID"`
	TiketID    int         `json:"tiket_id" gorm:"constrain:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Tiket      TiketRespon `json:"tiket" gorm:"constrain:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt  time.Time   `json:"-"`
	UpdatedAt  time.Time   `json:"-"`
}

type TransaksiResponse struct {
	ID         int         `json:"id"`
	Qty        int         `json:"qty"`
	Total      int         `json:"total"`
	Status     string      `json:"status"`
	Attachment string      `json:"attachment"`
	UserID     int         `json:"user_id"`
	User       UserRespon  `json:"user" gorm:"foreignKey:UserID"`
	TiketID    int         `json:"tiket_id" gorm:"constrain:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Tiket      TiketRespon `json:"tiket" gorm:"constrain:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (TransaksiResponse) TableName() string {
	return "transaksis"
}

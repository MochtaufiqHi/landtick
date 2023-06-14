package models

import "time"

type Tiket struct {
	ID           int           `json:"id" gorm:"primary_key:auto_increment"`
	Name         string        `json:"name" form:"name" gorm:"type: varchar(255)"`
	TrainID      int           `json:"train_id" gorm:"constrain:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Train        TrainResponse `json:"train" gorm:"constrain:OnUpdate:CASCADE,OnDelete:CASCADE"`
	JamBerangkat string        `json:"jam_berangkat" form:"jam_berangkat" gorm:"type: varchar(255)"`
	JamTiba      string        `json:"jam_tiba" form:"jam_tiba" gorm:"type: varchar(255)"`
	Durasi       string        `json:"durasi" form:"durasi" gorm:"type: varchar(255)"`
	Harga        int           `json:"harga" form:"harga" gorm:"type: int"`
	Tanggal      string        `json:"tanggal" form:"tanggal" gorm:"type: varchar(255)"`
	Kuota        int           `json:"kuota" form:"kuota" gorm:"type: int"`
	StasiunAwal  string        `json:"stasiun_awal" form:"stasiun_awal" gorm:"type: varchar(255)"`
	StasiunAkhir string        `json:"stasiun_akhir" form:"stasiun_akhir" gorm:"type: varchar(255)"`
	CreatedAt    time.Time     `json:"-"`
	UpdatedAt    time.Time     `json:"-"`
}

type TiketRespon struct {
	ID           int           `json:"id"`
	Name         string        `json:"name" `
	TrainID      int           `json:"train_id" gorm:"constrain:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Train        TrainResponse `json:"train" gorm:"foreignkey:TrainID"`
	JamBerangkat string        `json:"jam_berangkat"`
	JamTiba      string        `json:"jam_tiba"`
	Durasi       string        `json:"durasi"`
	Harga        int           `json:"harga"`
	Tanggal      string        `json:"tanggal"`
	Kuota        int           `json:"kuota"`
	StasiunAwal  string        `json:"stasiun_awal"`
	StasiunAkhir string        `json:"stasiun_akhir"`
	CreatedAt    time.Time     `json:"-"`
	UpdatedAt    time.Time     `json:"-"`
}

func (TiketRespon) TableName() string {
	return "tikets"
}

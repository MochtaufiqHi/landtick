package tiketdto

import "time"

type TiketRespon struct {
	ID      int    `json:"id"`
	Name    string `json:"name" `
	TrainID int    `json:"train_id" gorm:"constrain:OnUpdate:CASCADE,OnDelete:CASCADE"`
	// Train        TrainResponse `json:"train" gorm:"foreignkey:TrainID"`
	JamBerangkat string    `json:"jam_berangkat"`
	JamTiba      string    `json:"jam_tiba"`
	Durasi       string    `json:"durasi"`
	Harga        int       `json:"harga"`
	Tanggal      string    `json:"tanggal"`
	Kuota        int       `json:"kuota"`
	StasiunAwal  string    `json:"stasiun_awal"`
	StasiunAkhir string    `json:"stasiun_akhir"`
	CreatedAt    time.Time `json:"-"`
	UpdatedAt    time.Time `json:"-"`
}

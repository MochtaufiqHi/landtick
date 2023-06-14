package tiketdto

type TiketRequest struct {
	Name    string `json:"nama"`
	TrainID int    `json:"train_id"`
	// Train        models.TrainResponse `json:"train_kereta"`
	JamBerangkat string `json:"jam_berangkat"`
	JamTiba      string `json:"jam_tiba"`
	StasiunAwal  string `json:"stasiun_awal"`
	StasiunAkhir string `json:"stasiun_Akhir"`
	Durasi       string `json:"durasi"`
	Harga        int    `json:"harga"`
	Tanggal      string `json:"tanggal"`
	Kuota        int    `json:"kuota"`
}

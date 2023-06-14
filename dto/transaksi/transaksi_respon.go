package transaksidto

type TransaksiRespon struct {
	ID         int    `json:"id"`
	Qty        int    `json:"qty"`
	Total      int    `json:"total"`
	Status     string `json:"status"`
	Attachment string `json:"attachment"`
	UserID     int    `json:"user_id"`
	TiketID    int    `json:"tiket_id"`
}

package models

type Train struct {
	ID   int    `json:"id"`
	Name string `json:"name" gorm:"type: varchar(255)"`
}

type TrainResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (TrainResponse) TableName() string {
	return "trains"
}

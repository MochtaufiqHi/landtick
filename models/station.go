package models

type Station struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
	Kota string `json:"kota"`
}

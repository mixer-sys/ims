package models

type Product struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	CategoryID int     `json:"category_id"`
	Price      float64 `json:"price"`
}

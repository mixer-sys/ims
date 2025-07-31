package models

type Product struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	CategoryID string  `json:"category_id"`
	Price      float64 `json:"price"`
}

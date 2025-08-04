package models

import "github.com/go-playground/validator/v10"

type Product struct {
	ID         string  `json:"id" validate:"required,uuid"`
	Name       string  `json:"name" validate:"required"`
	CategoryID string  `json:"category_id"`
	Price      float64 `json:"price" validate:"required,gt=0"`
}

func ValidateProduct(product *Product) error {
	validate := validator.New()
	return validate.Struct(product)
}

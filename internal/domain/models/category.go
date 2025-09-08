package models

import "github.com/go-playground/validator/v10"

type Category struct {
	ID   string `json:"id" validate:"required,uuid"`
	Name string `json:"name" validate:"required"`
}

func ValidateCategory(category *Category) error {
	validate := validator.New()
	return validate.Struct(category)
}

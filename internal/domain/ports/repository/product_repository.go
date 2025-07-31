package repository

import (
	"context"
	"ims/internal/domain/models"
)

type ProductRepository interface {
	Create(ctx context.Context, product *models.Product) error
	GetByID(ctx context.Context, id string) (*models.Product, error)
	Update(ctx context.Context, product *models.Product) error
	Delete(ctx context.Context, id string) error
	GetAll(ctx context.Context, limit, offset int) ([]models.Product, error)
}

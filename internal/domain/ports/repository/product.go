package repository

import (
	"context"
	"fmt"
	"ims/internal/domain/models"

	"github.com/jackc/pgx/v4/pgxpool"
)

type productRepository struct {
	db *pgxpool.Pool
}

func NewProductRepository(db *pgxpool.Pool) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) Create(ctx context.Context, product *models.Product) error {
	query := "INSERT INTO products (name, category_id, price) VALUES ($1, $2, $3) RETURNING id"
	return r.db.QueryRow(
		ctx, query, product.Name, product.CategoryID, product.Price).Scan(&product.ID)
}

func (r *productRepository) GetByID(ctx context.Context, id int) (*models.Product, error) {
	var product models.Product
	query := "SELECT id, name, category_id, price FROM products WHERE id = $1"

	err := r.db.QueryRow(ctx, query, id).Scan(
		&product.ID, &product.Name, &product.CategoryID, &product.Price)
	if err != nil {
		return nil, fmt.Errorf("product not found: %w", err)
	}
	return &product, nil
}

func (r *productRepository) Update(ctx context.Context, product *models.Product) error {
	query := "UPDATE products SET name = $1, category_id = $2, price = $3 WHERE id = $4"
	_, err := r.db.Exec(
		ctx, query, product.Name, product.CategoryID, product.Price, product.ID)
	return fmt.Errorf("failed to update product: %w", err)
}

func (r *productRepository) Delete(ctx context.Context, id int) error {
	query := "DELETE FROM products WHERE id = $1"
	_, err := r.db.Exec(ctx, query, id)
	return fmt.Errorf("failed to delete product: %w", err)
}

func (r *productRepository) GetAll(ctx context.Context) ([]models.Product, error) {
	var products []models.Product
	query := "SELECT id, name, category_id, price FROM products"
	rows, err := r.db.Query(ctx, query)

	if err != nil {
		return nil, fmt.Errorf("failed to get products: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var product models.Product
		if err := rows.Scan(
			&product.ID, &product.Name, &product.CategoryID, &product.Price); err != nil {
			return nil, fmt.Errorf("failed to scan product: %w", err)
		}
		products = append(products, product)
	}
	return products, nil
}

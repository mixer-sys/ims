package repository

import (
	"context"
	"fmt"
	"ims/internal/domain/models"
	"ims/internal/interfaces/http/handlers"

	"github.com/jackc/pgx/v4/pgxpool"
)

type SQLProductRepository struct {
	db *pgxpool.Pool
}

func NewProductRepository(db *pgxpool.Pool) handlers.ProductRepository {
	return &SQLProductRepository{db: db}
}

func (r *SQLProductRepository) Create(ctx context.Context, product *models.Product) error {
	query := "INSERT INTO products (name, category_id, price) VALUES ($1, $2, $3) RETURNING id"

	return r.db.QueryRow(
		ctx, query, product.Name, product.CategoryID, product.Price).Scan(&product.ID)
}

func (r *SQLProductRepository) GetByID(ctx context.Context, id string) (*models.Product, error) {
	var product models.Product

	query := "SELECT id, name, category_id, price FROM products WHERE id = $1"

	err := r.db.QueryRow(ctx, query, id).Scan(
		&product.ID, &product.Name, &product.CategoryID, &product.Price)
	if err != nil {
		return nil, fmt.Errorf("product not found: %w", err)
	}

	return &product, nil
}

func (r *SQLProductRepository) Update(ctx context.Context, product *models.Product) error {
	query := "UPDATE products SET name = $1, category_id = $2, price = $3 WHERE id = $4"
	_, err := r.db.Exec(
		ctx, query, product.Name, product.CategoryID, product.Price, product.ID)

	if err != nil {
		return fmt.Errorf("failed to update product: %w", err)
	}

	return fmt.Errorf("failed to update product: %w", err)
}

func (r *SQLProductRepository) Delete(ctx context.Context, id string) error {
	query := "DELETE FROM products WHERE id = $1"
	_, err := r.db.Exec(ctx, query, id)

	if err != nil {
		return fmt.Errorf("failed to delete product: %w", err)
	}

	return fmt.Errorf("failed to delete product: %w", err)
}

func (r *SQLProductRepository) GetAll(ctx context.Context, limit, offset int) ([]models.Product, error) {
	var products []models.Product

	query := "SELECT id, name, category_id, price FROM products LIMIT $1 OFFSET $2"

	rows, err := r.db.Query(ctx, query, limit, offset)

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

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error occurred during rows iteration: %w", err)
	}

	return products, nil
}

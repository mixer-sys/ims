package repository

import (
	"context"
	"fmt"
	"ims/internal/domain/models"

	"github.com/jackc/pgx/v4/pgxpool"
)

type categoryRepository struct {
	db *pgxpool.Pool
}

func NewCategoryRepository(db *pgxpool.Pool) CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) Create(ctx context.Context, category *models.Category) error {
	query := "INSERT INTO categories (name) VALUES ($1) RETURNING id"

	return r.db.QueryRow(ctx, query, category.Name).Scan(&category.ID)
}

func (r *categoryRepository) GetByID(ctx context.Context, id int) (*models.Category, error) {
	var category models.Category

	query := "SELECT id, name FROM categories WHERE id = $1"

	err := r.db.QueryRow(ctx, query, id).Scan(&category.ID, &category.Name)
	if err != nil {
		return nil, fmt.Errorf("category not found: %w", err)
	}
	return &category, nil
}

func (r *categoryRepository) Update(ctx context.Context, category *models.Category) error {
	query := "UPDATE categories SET name = $1 WHERE id = $2"
	_, err := r.db.Exec(ctx, query, category.Name, category.ID)
	return fmt.Errorf("failed to update category: %w", err)
}

func (r *categoryRepository) Delete(ctx context.Context, id int) error {
	query := "DELETE FROM categories WHERE id = $1"
	_, err := r.db.Exec(ctx, query, id)
	return fmt.Errorf("failed to delete category: %w", err)
}

func (r *categoryRepository) GetAll(ctx context.Context) ([]models.Category, error) {
	var categories []models.Category
	query := "SELECT id, name FROM categories"
	rows, err := r.db.Query(ctx, query)

	if err != nil {
		return nil, fmt.Errorf("failed to get categories: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var category models.Category
		if err := rows.Scan(&category.ID, &category.Name); err != nil {
			return nil, fmt.Errorf("failed to scan category: %w", err)
		}
		categories = append(categories, category)
	}
	return categories, nil
}

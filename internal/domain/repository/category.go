package repository

import (
	"context"
	"errors"
	"fmt"
	"ims/internal/domain/handlers"
	"ims/internal/domain/models"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v4/pgxpool"
)

type SQLCategoryRepository struct {
	db *pgxpool.Pool
}

func NewCategoryRepository(db *pgxpool.Pool) handlers.CategoryRepository {
	return &SQLCategoryRepository{db: db}
}

func (r *SQLCategoryRepository) Create(ctx context.Context, category *models.Category) error {
	if err := models.ValidateCategory(category); err != nil {
		return fmt.Errorf("validation error: %w", err)
	}
	query := "INSERT INTO categories (name) VALUES ($1) RETURNING id"

	return r.db.QueryRow(ctx, query, category.Name).Scan(&category.ID)
}

func (r *SQLCategoryRepository) GetByID(ctx context.Context, id string) (*models.Category, error) {
	var category models.Category

	query := "SELECT id, name FROM categories WHERE id = $1"

	err := r.db.QueryRow(ctx, query, id).Scan(&category.ID, &category.Name)
	if err == pgx.ErrNoRows {
		return nil, fmt.Errorf("category not found: %w", err)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get category: %w", err)
	}

	return &category, nil
}

func (r *SQLCategoryRepository) Update(ctx context.Context, category *models.Category) error {
	if err := models.ValidateCategory(category); err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	query := "UPDATE categories SET name = $1 WHERE id = $2"
	_, err := r.db.Exec(ctx, query, category.Name, category.ID)
	if err != nil {
		return fmt.Errorf("failed to update category: %w", err)
	}

	return nil
}

func (r *SQLCategoryRepository) Delete(ctx context.Context, id string) error {
	query := "DELETE FROM categories WHERE id = $1"
	_, err := r.db.Exec(ctx, query, id)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return fmt.Errorf("category not found: %w", err)
		}
		return fmt.Errorf("failed to delete category: %w", err)
	}

	return nil
}

func (r *SQLCategoryRepository) GetAll(ctx context.Context, limit, offset int) ([]models.Category, error) {
	categories := make([]models.Category, 0)

	query := "SELECT id, name FROM categories LIMIT $1 OFFSET $2"

	rows, err := r.db.Query(ctx, query, limit, offset)

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

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error occurred during rows iteration: %w", err)
	}

	return categories, nil
}

package repository

import (
	"database/sql"
	"ims/internal/domain/models"
)

type CategoryRepository interface {
	Create(category *models.Category) error
	GetByID(id int) (*models.Category, error)
	Update(category *models.Category) error
	Delete(id int) error
	GetAll() ([]models.Category, error)
}

type categoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) Create(category *models.Category) error {
	query := "INSERT INTO categories (name) VALUES ($1) RETURNING id"
	return r.db.QueryRow(query, category.Name).Scan(&category.ID)
}

func (r *categoryRepository) GetByID(id int) (*models.Category, error) {
	var category models.Category
	query := "SELECT id, name FROM categories WHERE id = $1"
	err := r.db.QueryRow(query, id).Scan(&category.ID, &category.Name)
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *categoryRepository) Update(category *models.Category) error {
	query := "UPDATE categories SET name = $1 WHERE id = $2"
	_, err := r.db.Exec(query, category.Name, category.ID)
	return err
}

func (r *categoryRepository) Delete(id int) error {
	query := "DELETE FROM categories WHERE id = $1"
	_, err := r.db.Exec(query, id)
	return err
}

func (r *categoryRepository) GetAll() ([]models.Category, error) {
	var categories []models.Category
	query := "SELECT id, name FROM categories"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var category models.Category
		if err := rows.Scan(&category.ID, &category.Name); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}

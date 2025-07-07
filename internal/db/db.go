package db

import (
	"database/sql"
	"ims/internal/models"
)

type ProductRepository interface {
	Create(product *models.Product) error
	GetByID(id int) (*models.Product, error)
	Update(product *models.Product) error
	Delete(id int) error
	GetAll() ([]models.Product, error)
}

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepository{db: db}
}

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

func (r *productRepository) Create(product *models.Product) error {
	query := "INSERT INTO products (name, category_id, price) VALUES ($1, $2, $3) RETURNING id"
	return r.db.QueryRow(query, product.Name, product.CategoryID, product.Price).Scan(&product.ID)
}

func (r *productRepository) GetByID(id int) (*models.Product, error) {
	var product models.Product
	query := "SELECT id, name, category_id, price FROM products WHERE id = $1"
	err := r.db.QueryRow(query, id).Scan(&product.ID, &product.Name, &product.CategoryID, &product.Price)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productRepository) Update(product *models.Product) error {
	query := "UPDATE products SET name = $1, category_id = $2, price = $3 WHERE id = $4"
	_, err := r.db.Exec(query, product.Name, product.CategoryID, product.Price, product.ID)
	return err
}

func (r *productRepository) Delete(id int) error {
	query := "DELETE FROM products WHERE id = $1"
	_, err := r.db.Exec(query, id)
	return err
}

func (r *productRepository) GetAll() ([]models.Product, error) {
	var products []models.Product
	query := "SELECT id, name, category_id, price FROM products"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.CategoryID, &product.Price); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
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

package db

import (
	"ims/internal/db"
	"ims/internal/models"
)

type ProductService interface {
	Create(product *models.Product) error
	GetByID(id int) (*models.Product, error)
	Update(product *models.Product) error
	Delete(id int) error
	GetAll() ([]models.Product, error)
}

type productService struct {
	repo db.ProductRepository
}

func NewProductService(repo db.ProductRepository) ProductService {
	return &productService{repo: repo}
}

type CategoryService interface {
	Create(category *models.Category) error
	GetByID(id int) (*models.Category, error)
	Update(category *models.Category) error
	Delete(id int) error
	GetAll() ([]models.Category, error)
}

type categoryService struct {
	repo db.CategoryRepository
}

func NewCategoryService(repo db.CategoryRepository) CategoryService {
	return &categoryService{repo: repo}
}

func (s *productService) Create(product *models.Product) error {
	return s.repo.Create(product)
}

func (s *productService) GetByID(id int) (*models.Product, error) {
	return s.repo.GetByID(id)
}

func (s *productService) Update(product *models.Product) error {
	return s.repo.Update(product)
}

func (s *productService) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *productService) GetAll() ([]models.Product, error) {
	return s.repo.GetAll()
}

func (s *categoryService) Create(category *models.Category) error {
	return s.repo.Create(category)
}

func (s *categoryService) GetByID(id int) (*models.Category, error) {
	return s.repo.GetByID(id)
}

func (s *categoryService) Update(category *models.Category) error {
	return s.repo.Update(category)
}

func (s *categoryService) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *categoryService) GetAll() ([]models.Category, error) {
	return s.repo.GetAll()
}

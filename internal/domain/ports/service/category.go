package service

import (
	"ims/internal/domain/models"
	"ims/internal/domain/ports/repository"
)

type CategoryService interface {
	Create(category *models.Category) error
	GetByID(id int) (*models.Category, error)
	Update(category *models.Category) error
	Delete(id int) error
	GetAll() ([]models.Category, error)
}

type categoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) CategoryService {
	return &categoryService{repo: repo}
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

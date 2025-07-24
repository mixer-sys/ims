package service

import (
	"context"
	"ims/internal/domain/models"
	"ims/internal/domain/ports/repository"
)

type CategoryService interface {
	Create(ctx context.Context, category *models.Category) error
	GetByID(ctx context.Context, id int) (*models.Category, error)
	Update(ctx context.Context, category *models.Category) error
	Delete(ctx context.Context, id int) error
	GetAll(ctx context.Context) ([]models.Category, error)
}

type categoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) CategoryService {
	return &categoryService{repo: repo}
}

func (s *categoryService) Create(ctx context.Context, category *models.Category) error {
	return s.repo.Create(ctx, category)
}

func (s *categoryService) GetByID(ctx context.Context, id int) (*models.Category, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *categoryService) Update(ctx context.Context, category *models.Category) error {
	return s.repo.Update(ctx, category)
}

func (s *categoryService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

func (s *categoryService) GetAll(ctx context.Context) ([]models.Category, error) {
	return s.repo.GetAll(ctx)
}

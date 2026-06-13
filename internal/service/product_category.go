package service

import (
	"fmt"
	"procurement-system/internal/model"
	"procurement-system/internal/repository"
)

type ProductCategoryService struct {
	repo repository.ProductCategoryRepository
}

func NewProductCategoryService(repo repository.ProductCategoryRepository) *ProductCategoryService {
	return &ProductCategoryService{repo: repo}
}

func (s *ProductCategoryService) GetAll(filter model.ListFilter) ([]model.ProductCategory, int, error) {
	filter.Normalize()
	return s.repo.GetAll(filter)
}

func (s *ProductCategoryService) GetByID(id int) (*model.ProductCategory, error) {
	if id <= 0 {
		return nil, fmt.Errorf("invalid category ID")
	}
	return s.repo.GetByID(id)
}

func (s *ProductCategoryService) Create(pc *model.ProductCategory) error {
	if pc.Name == "" {
		return fmt.Errorf("category name is required")
	}
	return s.repo.Create(pc)
}

func (s *ProductCategoryService) Update(pc *model.ProductCategory) error {
	if pc.ID <= 0 {
		return fmt.Errorf("invalid category ID")
	}
	return s.repo.Update(pc)
}

func (s *ProductCategoryService) Delete(id int) error {
	if id <= 0 {
		return fmt.Errorf("invalid category ID")
	}
	return s.repo.Delete(id)
}
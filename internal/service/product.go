package service

import (
	"fmt"
	"procurement-system/internal/model"
	"procurement-system/internal/repository"
)

type ProductService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetAll(filter model.ListFilter) ([]model.Product, int, error) {
	filter.Normalize()
	return s.repo.GetAll(filter)
}

func (s *ProductService) GetByID(id int) (*model.Product, error) {
	if id <= 0 {
		return nil, fmt.Errorf("invalid product ID")
	}
	return s.repo.GetByID(id)
}

func (s *ProductService) Create(p *model.Product) error {
	if p.Name == "" || p.SKU == "" {
		return fmt.Errorf("product name and SKU are required")
	}
	return s.repo.Create(p)
}

func (s *ProductService) Update(p *model.Product) error {
	if p.ID <= 0 {
		return fmt.Errorf("invalid product ID")
	}
	if p.Name == "" || p.SKU == "" {
		return fmt.Errorf("product name and SKU are required")
	}
	return s.repo.Update(p)
}

func (s *ProductService) Delete(id int) error {
	if id <= 0 {
		return fmt.Errorf("invalid product ID")
	}
	return s.repo.Delete(id)
}
package service

import (
	"fmt"
	"procurement-system/internal/model"
	"procurement-system/internal/repository"
)

type ProductUnitService struct {
	repo repository.ProductUnitRepository
}

func NewProductUnitService(repo repository.ProductUnitRepository) *ProductUnitService {
	return &ProductUnitService{repo: repo}
}

func (s *ProductUnitService) GetAll(filter model.ListFilter) ([]model.ProductUnit, int, error) {
	filter.Normalize()
	return s.repo.GetAll(filter)
}

func (s *ProductUnitService) GetByID(id int) (*model.ProductUnit, error) {
	if id <= 0 {
		return nil, fmt.Errorf("invalid unit ID")
	}
	return s.repo.GetByID(id)
}

func (s *ProductUnitService) Create(pu *model.ProductUnit) error {
	if pu.Name == "" || pu.ShortName == "" {
		return fmt.Errorf("unit name and short name are required")
	}
	return s.repo.Create(pu)
}

func (s *ProductUnitService) Update(pu *model.ProductUnit) error {
	if pu.ID <= 0 {
		return fmt.Errorf("invalid unit ID")
	}
	return s.repo.Update(pu)
}

func (s *ProductUnitService) Delete(id int) error {
	if id <= 0 {
		return fmt.Errorf("invalid unit ID")
	}
	return s.repo.Delete(id)
}
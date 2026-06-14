package service

import (
	"fmt"
	"procurement-system/internal/model"
	"procurement-system/internal/repository"
)

type SupplierService struct {
	repo repository.SupplierRepository
}

func NewSupplierService(repo repository.SupplierRepository) *SupplierService {
	return &SupplierService{repo: repo}
}

func (s *SupplierService) GetAll(filter model.ListFilter) ([]model.Supplier, int, error) {
	filter.Normalize()
	return s.repo.GetAll(filter)
}

func (s *SupplierService) GetByID(id int) (*model.Supplier, error) {
	if id <= 0 {
		return nil, fmt.Errorf("invalid supplier ID")
	}
	return s.repo.GetByID(id)
}

func (s *SupplierService) Create(sup *model.Supplier) error {
	if sup.Name == "" {
		return fmt.Errorf("supplier name is required")
	}
	return s.repo.Create(sup)
}

func (s *SupplierService) Update(sup *model.Supplier) error {
	if sup.ID <= 0 {
		return fmt.Errorf("invalid supplier ID")
	}
	if sup.Name == "" {
		return fmt.Errorf("supplier name is required")
	}
	return s.repo.Update(sup)
}

func (s *SupplierService) Delete(id int) error {
	if id <= 0 {
		return fmt.Errorf("invalid supplier ID")
	}
	return s.repo.Delete(id)
}
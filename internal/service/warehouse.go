package service

import (
	"fmt"
	"procurement-system/internal/model"
	"procurement-system/internal/repository"
)

type WarehouseService struct {
	repo repository.WarehouseRepository
}

func NewWarehouseService(repo repository.WarehouseRepository) *WarehouseService {
	return &WarehouseService{repo: repo}
}

func (s *WarehouseService) GetAll(filter model.ListFilter) ([]model.Warehouse, int, error) {
	filter.Normalize()
	return s.repo.GetAll(filter)
}

func (s *WarehouseService) GetByID(id int) (*model.Warehouse, error) {
	if id <= 0 {
		return nil, fmt.Errorf("invalid warehouse ID")
	}
	return s.repo.GetByID(id)
}

func (s *WarehouseService) Create(w *model.Warehouse) error {
	if w.WarehouseType == "" {
		return fmt.Errorf("warehouse type is required")
	}
	return s.repo.Create(w)
}

func (s *WarehouseService) Update(w *model.Warehouse) error {
	if w.ID <= 0 {
		return fmt.Errorf("invalid warehouse ID")
	}
	return s.repo.Update(w)
}

func (s *WarehouseService) Delete(id int) error {
	if id <= 0 {
		return fmt.Errorf("invalid warehouse ID")
	}
	return s.repo.Delete(id)
}
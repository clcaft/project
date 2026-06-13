package service

import (
	"procurement-system/internal/model"
	"procurement-system/internal/repository"
)

type InventoryBalanceService struct {
	repo repository.InventoryBalanceRepository
}

func NewInventoryBalanceService(repo repository.InventoryBalanceRepository) *InventoryBalanceService {
	return &InventoryBalanceService{repo: repo}
}

func (s *InventoryBalanceService) GetAll(filter model.ListFilter) ([]model.InventoryBalance, int, error) {
	filter.Normalize()
	return s.repo.GetAll(filter)
}

func (s *InventoryBalanceService) GetByDepartment(departmentID int) ([]model.InventoryBalance, error) {
	return s.repo.GetByDepartment(departmentID)
}
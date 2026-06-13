package service

import (
	"procurement-system/internal/model"
	"procurement-system/internal/repository"
)

type ReportService struct {
	repos *repository.Repositories
}

func NewReportService(repos *repository.Repositories) *ReportService {
	return &ReportService{repos: repos}
}

func (s *ReportService) GetInventoryReport(departmentID int) ([]model.InventoryBalance, error) {
	if departmentID > 0 {
		return s.repos.InventoryBalance.GetByDepartment(departmentID)
	}
	balances, _, err := s.repos.InventoryBalance.GetAll(model.ListFilter{Page: 1, PerPage: 1000})
	return balances, err
}
package service

import (
	"fmt"
	"procurement-system/internal/model"
	"procurement-system/internal/repository"
)

type DepartmentService struct {
	repo repository.DepartmentRepository
}

func NewDepartmentService(repo repository.DepartmentRepository) *DepartmentService {
	return &DepartmentService{repo: repo}
}

func (s *DepartmentService) GetAll(filter model.ListFilter) ([]model.Department, int, error) {
	filter.Normalize()
	return s.repo.GetAll(filter)
}

func (s *DepartmentService) GetByID(id int) (*model.Department, error) {
	if id <= 0 {
		return nil, fmt.Errorf("invalid department ID")
	}
	return s.repo.GetByID(id)
}

func (s *DepartmentService) Create(d *model.Department) error {
	if d.Name == "" {
		return fmt.Errorf("department name is required")
	}
	return s.repo.Create(d)
}

func (s *DepartmentService) Update(d *model.Department) error {
	if d.ID <= 0 {
		return fmt.Errorf("invalid department ID")
	}
	if d.Name == "" {
		return fmt.Errorf("department name is required")
	}
	return s.repo.Update(d)
}

func (s *DepartmentService) Delete(id int) error {
	if id <= 0 {
		return fmt.Errorf("invalid department ID")
	}
	return s.repo.Delete(id)
}
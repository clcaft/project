package service

import (
	"fmt"
	"procurement-system/internal/model"
	"procurement-system/internal/repository"
)

type StoreService struct {
	repo repository.StoreRepository
}

func NewStoreService(repo repository.StoreRepository) *StoreService {
	return &StoreService{repo: repo}
}

func (s *StoreService) GetAll(filter model.ListFilter) ([]model.Store, int, error) {
	filter.Normalize()
	return s.repo.GetAll(filter)
}

func (s *StoreService) GetByID(id int) (*model.Store, error) {
	if id <= 0 {
		return nil, fmt.Errorf("invalid store ID")
	}
	return s.repo.GetByID(id)
}

func (s *StoreService) Create(st *model.Store) error {
	if st.StoreType == "" {
		return fmt.Errorf("store type is required")
	}
	return s.repo.Create(st)
}

func (s *StoreService) Update(st *model.Store) error {
	if st.ID <= 0 {
		return fmt.Errorf("invalid store ID")
	}
	return s.repo.Update(st)
}

func (s *StoreService) Delete(id int) error {
	if id <= 0 {
		return fmt.Errorf("invalid store ID")
	}
	return s.repo.Delete(id)
}
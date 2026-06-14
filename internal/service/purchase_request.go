package service

import (
	"fmt"
	"procurement-system/internal/model"
	"procurement-system/internal/repository"
)

type PurchaseRequestService struct {
	repo repository.PurchaseRequestRepository
}

func NewPurchaseRequestService(repo repository.PurchaseRequestRepository) *PurchaseRequestService {
	return &PurchaseRequestService{repo: repo}
}

func (s *PurchaseRequestService) GetAll(filter model.ListFilter) ([]model.PurchaseRequest, int, error) {
	filter.Normalize()
	return s.repo.GetAll(filter)
}

func (s *PurchaseRequestService) GetByID(id int) (*model.PurchaseRequest, error) {
	if id <= 0 {
		return nil, fmt.Errorf("invalid purchase request ID")
	}
	return s.repo.GetByID(id)
}

func (s *PurchaseRequestService) Create(pr *model.PurchaseRequest) error {
	if pr.RequestDate == "" {
		return fmt.Errorf("request date is required")
	}
	if pr.RecipientDepartmentID <= 0 {
		return fmt.Errorf("recipient department is required")
	}
	if len(pr.Items) == 0 {
		return fmt.Errorf("at least one item is required")
	}
	pr.Status = "draft"
	return s.repo.Create(pr)
}

func (s *PurchaseRequestService) Update(pr *model.PurchaseRequest) error {
	if pr.ID <= 0 {
		return fmt.Errorf("invalid purchase request ID")
	}
	return s.repo.Update(pr)
}

func (s *PurchaseRequestService) Delete(id int) error {
	if id <= 0 {
		return fmt.Errorf("invalid purchase request ID")
	}
	return s.repo.Delete(id)
}

func (s *PurchaseRequestService) Submit(id int) error {
	return s.repo.ChangeStatus(id, "submitted")
}

func (s *PurchaseRequestService) Approve(id int) error {
	return s.repo.ChangeStatus(id, "approved")
}

func (s *PurchaseRequestService) Reject(id int) error {
	return s.repo.ChangeStatus(id, "rejected")
}

func (s *PurchaseRequestService) Cancel(id int) error {
	return s.repo.ChangeStatus(id, "cancelled")
}
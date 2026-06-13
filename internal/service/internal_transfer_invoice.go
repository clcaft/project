package service

import (
	"fmt"
	"procurement-system/internal/model"
	"procurement-system/internal/repository"
)

type InternalTransferInvoiceService struct {
	repo repository.InternalTransferInvoiceRepository
}

func NewInternalTransferInvoiceService(repo repository.InternalTransferInvoiceRepository) *InternalTransferInvoiceService {
	return &InternalTransferInvoiceService{repo: repo}
}

func (s *InternalTransferInvoiceService) GetAll(filter model.ListFilter) ([]model.InternalTransferInvoice, int, error) {
	filter.Normalize()
	return s.repo.GetAll(filter)
}

func (s *InternalTransferInvoiceService) GetByID(id int) (*model.InternalTransferInvoice, error) {
	if id <= 0 {
		return nil, fmt.Errorf("invalid transfer invoice ID")
	}
	return s.repo.GetByID(id)
}

func (s *InternalTransferInvoiceService) Create(it *model.InternalTransferInvoice) error {
	if it.TransferDate == "" {
		return fmt.Errorf("transfer date is required")
	}
	if it.FromDepartmentID <= 0 || it.ToDepartmentID <= 0 {
		return fmt.Errorf("both departments are required")
	}
	if it.FromDepartmentID == it.ToDepartmentID {
		return fmt.Errorf("source and destination departments must be different")
	}
	if len(it.Items) == 0 {
		return fmt.Errorf("at least one item is required")
	}
	it.Status = "draft"
	return s.repo.Create(it)
}

func (s *InternalTransferInvoiceService) Update(it *model.InternalTransferInvoice) error {
	if it.ID <= 0 {
		return fmt.Errorf("invalid transfer invoice ID")
	}
	return s.repo.Update(it)
}

func (s *InternalTransferInvoiceService) Delete(id int) error {
	if id <= 0 {
		return fmt.Errorf("invalid transfer invoice ID")
	}
	return s.repo.Delete(id)
}

func (s *InternalTransferInvoiceService) Confirm(id int) error {
	return s.repo.ChangeStatus(id, "confirmed")
}

func (s *InternalTransferInvoiceService) Ship(id int) error {
	return s.repo.ChangeStatus(id, "shipped")
}

func (s *InternalTransferInvoiceService) Receive(id int) error {
	return s.repo.ChangeStatus(id, "received")
}

func (s *InternalTransferInvoiceService) Cancel(id int) error {
	return s.repo.ChangeStatus(id, "cancelled")
}
package service

import (
	"fmt"
	"procurement-system/internal/model"
	"procurement-system/internal/repository"
)

type DeliveryInvoiceService struct {
	repo repository.DeliveryInvoiceRepository
}

func NewDeliveryInvoiceService(repo repository.DeliveryInvoiceRepository) *DeliveryInvoiceService {
	return &DeliveryInvoiceService{repo: repo}
}

func (s *DeliveryInvoiceService) GetAll(filter model.ListFilter) ([]model.DeliveryInvoice, int, error) {
	filter.Normalize()
	return s.repo.GetAll(filter)
}

func (s *DeliveryInvoiceService) GetByID(id int) (*model.DeliveryInvoice, error) {
	if id <= 0 {
		return nil, fmt.Errorf("invalid delivery invoice ID")
	}
	return s.repo.GetByID(id)
}

func (s *DeliveryInvoiceService) Create(di *model.DeliveryInvoice) error {
	if di.InvoiceDate == "" {
		return fmt.Errorf("invoice date is required")
	}
	if di.SupplierID <= 0 {
		return fmt.Errorf("supplier is required")
	}
	if di.RecipientDepartmentID <= 0 {
		return fmt.Errorf("recipient department is required")
	}
	if len(di.Items) == 0 {
		return fmt.Errorf("at least one item is required")
	}
	di.Status = "draft"
	return s.repo.Create(di)
}

func (s *DeliveryInvoiceService) Update(di *model.DeliveryInvoice) error {
	if di.ID <= 0 {
		return fmt.Errorf("invalid delivery invoice ID")
	}
	return s.repo.Update(di)
}

func (s *DeliveryInvoiceService) Delete(id int) error {
	if id <= 0 {
		return fmt.Errorf("invalid delivery invoice ID")
	}
	return s.repo.Delete(id)
}

func (s *DeliveryInvoiceService) Confirm(id int) error {
	return s.repo.ChangeStatus(id, "confirmed")
}

func (s *DeliveryInvoiceService) Receive(id int) error {
	return s.repo.ChangeStatus(id, "received")
}

func (s *DeliveryInvoiceService) Cancel(id int) error {
	return s.repo.ChangeStatus(id, "cancelled")
}
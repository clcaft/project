package service

import "procurement-system/internal/repository"

type Services struct {
	Department              *DepartmentService
	Warehouse               *WarehouseService
	Store                   *StoreService
	Supplier                *SupplierService
	Product                 *ProductService
	ProductCategory         *ProductCategoryService
	ProductUnit             *ProductUnitService
	PurchaseRequest         *PurchaseRequestService
	DeliveryInvoice         *DeliveryInvoiceService
	InternalTransferInvoice *InternalTransferInvoiceService
	InventoryBalance        *InventoryBalanceService
	Report                  *ReportService
}

func NewServices(repos *repository.Repositories) *Services {
	return &Services{
		Department:              NewDepartmentService(repos.Department),
		Warehouse:               NewWarehouseService(repos.Warehouse),
		Store:                   NewStoreService(repos.Store),
		Supplier:                NewSupplierService(repos.Supplier),
		Product:                 NewProductService(repos.Product),
		ProductCategory:         NewProductCategoryService(repos.ProductCategory),
		ProductUnit:             NewProductUnitService(repos.ProductUnit),
		PurchaseRequest:         NewPurchaseRequestService(repos.PurchaseRequest),
		DeliveryInvoice:         NewDeliveryInvoiceService(repos.DeliveryInvoice),
		InternalTransferInvoice: NewInternalTransferInvoiceService(repos.InternalTransferInvoice),
		InventoryBalance:        NewInventoryBalanceService(repos.InventoryBalance),
		Report:                  NewReportService(repos),
	}
}
package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"procurement-system/internal/logger"
	"procurement-system/internal/model"
	"procurement-system/internal/service"
)

type Handlers struct {
	Department              *DepartmentHandler
	Warehouse               *WarehouseHandler
	Store                   *StoreHandler
	Supplier                *SupplierHandler
	Product                 *ProductHandler
	ProductCategory         *ProductCategoryHandler
	ProductUnit             *ProductUnitHandler
	PurchaseRequest         *PurchaseRequestHandler
	DeliveryInvoice         *DeliveryInvoiceHandler
	InternalTransferInvoice *InternalTransferInvoiceHandler
	InventoryBalance        *InventoryBalanceHandler
	Report                  *ReportHandler
}

func NewHandlers(services *service.Services, log logger.Logger) *Handlers {
	return &Handlers{
		Department:              NewDepartmentHandler(services.Department, log),
		Warehouse:               NewWarehouseHandler(services.Warehouse, log),
		Store:                   NewStoreHandler(services.Store, log),
		Supplier:                NewSupplierHandler(services.Supplier, log),
		Product:                 NewProductHandler(services.Product, log),
		ProductCategory:         NewProductCategoryHandler(services.ProductCategory, log),
		ProductUnit:             NewProductUnitHandler(services.ProductUnit, log),
		PurchaseRequest:         NewPurchaseRequestHandler(services.PurchaseRequest, log),
		DeliveryInvoice:         NewDeliveryInvoiceHandler(services.DeliveryInvoice, log),
		InternalTransferInvoice: NewInternalTransferInvoiceHandler(services.InternalTransferInvoice, log),
		InventoryBalance:        NewInventoryBalanceHandler(services.InventoryBalance, log),
		Report:                  NewReportHandler(services.Report, log),
	}
}

func respondJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func respondError(w http.ResponseWriter, status int, message string) {
	respondJSON(w, status, model.APIResponse{Success: false, Error: message})
}

func respondSuccess(w http.ResponseWriter, data interface{}) {
	respondJSON(w, http.StatusOK, model.APIResponse{Success: true, Data: data})
}

func respondCreated(w http.ResponseWriter, data interface{}) {
	respondJSON(w, http.StatusCreated, model.APIResponse{Success: true, Data: data})
}

func parseID(r *http.Request) (int, error) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		idStr = r.Context().Value("id").(string)
	}
	return strconv.Atoi(idStr)
}

func parseListFilter(r *http.Request) model.ListFilter {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	perPage, _ := strconv.Atoi(r.URL.Query().Get("per_page"))

	return model.ListFilter{
		Page:     page,
		PerPage:  perPage,
		Search:   r.URL.Query().Get("search"),
		SortBy:   r.URL.Query().Get("sort_by"),
		SortDir:  r.URL.Query().Get("sort_dir"),
		Status:   r.URL.Query().Get("status"),
		DateFrom: r.URL.Query().Get("date_from"),
		DateTo:   r.URL.Query().Get("date_to"),
	}
}
package repository

import (
	"database/sql"
	"fmt"
	"strings"

	"procurement-system/internal/model"
)

type Repositories struct {
	Department              DepartmentRepository
	Warehouse               WarehouseRepository
	Store                   StoreRepository
	Supplier                SupplierRepository
	Product                 ProductRepository
	ProductCategory         ProductCategoryRepository
	ProductUnit             ProductUnitRepository
	PurchaseRequest         PurchaseRequestRepository
	DeliveryInvoice         DeliveryInvoiceRepository
	InternalTransferInvoice InternalTransferInvoiceRepository
	InventoryBalance        InventoryBalanceRepository
}

func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		Department:              NewDepartmentRepository(db),
		Warehouse:               NewWarehouseRepository(db),
		Store:                   NewStoreRepository(db),
		Supplier:                NewSupplierRepository(db),
		Product:                 NewProductRepository(db),
		ProductCategory:         NewProductCategoryRepository(db),
		ProductUnit:             NewProductUnitRepository(db),
		PurchaseRequest:         NewPurchaseRequestRepository(db),
		DeliveryInvoice:         NewDeliveryInvoiceRepository(db),
		InternalTransferInvoice: NewInternalTransferInvoiceRepository(db),
		InventoryBalance:        NewInventoryBalanceRepository(db),
	}
}

func buildListQuery(baseQuery string, filter model.ListFilter, allowedSortFields map[string]string) (string, []interface{}) {
	var args []interface{}
	var conditions []string

	if filter.Search != "" {
		conditions = append(conditions, "name LIKE ?")
		args = append(args, "%"+filter.Search+"%")
	}
	if filter.Status != "" {
		conditions = append(conditions, "status = ?")
		args = append(args, filter.Status)
	}
	if filter.DateFrom != "" {
		conditions = append(conditions, "DATE(created_at) >= ?")
		args = append(args, filter.DateFrom)
	}
	if filter.DateTo != "" {
		conditions = append(conditions, "DATE(created_at) <= ?")
		args = append(args, filter.DateTo)
	}

	query := baseQuery
	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	if sortField, ok := allowedSortFields[filter.SortBy]; ok {
		query += fmt.Sprintf(" ORDER BY %s %s", sortField, filter.SortDir)
	} else {
		query += " ORDER BY id DESC"
	}

	offset := (filter.Page - 1) * filter.PerPage
	query += fmt.Sprintf(" LIMIT %d OFFSET %d", filter.PerPage, offset)

	return query, args
}

func countQuery(db *sql.DB, baseCountQuery string, filter model.ListFilter) (int, error) {
	var args []interface{}
	var conditions []string

	if filter.Search != "" {
		conditions = append(conditions, "name LIKE ?")
		args = append(args, "%"+filter.Search+"%")
	}
	if filter.Status != "" {
		conditions = append(conditions, "status = ?")
		args = append(args, filter.Status)
	}
	if filter.DateFrom != "" {
		conditions = append(conditions, "DATE(created_at) >= ?")
		args = append(args, filter.DateFrom)
	}
	if filter.DateTo != "" {
		conditions = append(conditions, "DATE(created_at) <= ?")
		args = append(args, filter.DateTo)
	}

	query := baseCountQuery
	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	var count int
	err := db.QueryRow(query, args...).Scan(&count)
	return count, err
}

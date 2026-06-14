package model

import (
	"database/sql"
	"time"
)

type Department struct {
	ID             int    `db:"id" json:"id"`
	Name           string `db:"name" json:"name"`
	DepartmentType string `db:"department_type" json:"department_type"`
	IsActive       bool   `db:"is_active" json:"is_active"`
}

type Warehouse struct {
	ID            int    `db:"id" json:"id"`
	Address       string `db:"address" json:"address"`
	WarehouseType string `db:"warehouse_type" json:"warehouse_type"`
	DepartmentID  int    `db:"department_id" json:"department_id"`
}

type Store struct {
	ID           int    `db:"id" json:"id"`
	Address      string `db:"address" json:"address"`
	StoreType    string `db:"store_type" json:"store_type"`
	DepartmentID int    `db:"department_id" json:"department_id"`
}

type Supplier struct {
	ID       int    `db:"id" json:"id"`
	Name     string `db:"name" json:"name"`
	Phone    string `db:"phone" json:"phone"`
	Email    string `db:"email" json:"email"`
	IsActive bool   `db:"is_active" json:"is_active"`
	Address  string `db:"address" json:"address"`
}

type ProductCategory struct {
	ID   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

type ProductUnit struct {
	ID        int    `db:"id" json:"id"`
	Name      string `db:"name" json:"name"`
	ShortName string `db:"short_name" json:"short_name"`
}

type Product struct {
	ID           int    `db:"id" json:"id"`
	Name         string `db:"name" json:"name"`
	SKU          string `db:"sku" json:"sku"`
	UnitID       int    `db:"unit_id" json:"unit_id"`
	CategoryID   *int   `db:"category_id" json:"category_id,omitempty"`
	IsActive     bool   `db:"is_active" json:"is_active"`
	UnitName     string `db:"unit_name" json:"unit_name,omitempty"`
	CategoryName string `db:"category_name" json:"category_name,omitempty"`
}

type SupplierProduct struct {
	ID         int `db:"id" json:"id"`
	SupplierID int `db:"supplier_id" json:"supplier_id"`
	ProductID  int `db:"product_id" json:"product_id"`
}

type PurchaseRequest struct {
	ID                      int       `db:"id" json:"id"`
	RequestDate             string    `db:"request_date" json:"request_date"`
	PlannedDeliveryDate     *string   `db:"planned_delivery_date" json:"planned_delivery_date,omitempty"`
	SupplierID              *int      `db:"supplier_id" json:"supplier_id,omitempty"`
	Status                  string    `db:"status" json:"status"`
	CreatedAt               time.Time `db:"created_at" json:"created_at"`
	CancelledAt             *time.Time `db:"cancelled_at" json:"cancelled_at,omitempty"`
	RecipientDepartmentID   int       `db:"recipient_department_id" json:"recipient_department_id"`
	SupplierName            string    `db:"supplier_name" json:"supplier_name,omitempty"`
	RecipientDepartmentName string    `db:"recipient_department_name" json:"recipient_department_name,omitempty"`
	Items                   []PurchaseRequestItem `json:"items,omitempty"`
}

type PurchaseRequestItem struct {
	ID                int     `db:"id" json:"id"`
	PurchaseRequestID int     `db:"purchase_request_id" json:"purchase_request_id"`
	ProductID         int     `db:"product_id" json:"product_id"`
	Quantity          float64 `db:"quantity" json:"quantity"`
	ProductName       string  `db:"product_name" json:"product_name,omitempty"`
	UnitName          string  `db:"unit_name" json:"unit_name,omitempty"`
}

type DeliveryInvoice struct {
	ID                      int            `db:"id" json:"id"`
	InvoiceDate             string         `db:"invoice_date" json:"invoice_date"`
	SupplierID              int            `db:"supplier_id" json:"supplier_id"`
	PurchaseRequestID       *int           `db:"purchase_request_id" json:"purchase_request_id,omitempty"`
	Status                  string         `db:"status" json:"status"`
	CreatedAt               time.Time      `db:"created_at" json:"created_at"`
	CancelledAt             *time.Time     `db:"cancelled_at" json:"cancelled_at,omitempty"`
	RecipientDepartmentID   int            `db:"recipient_department_id" json:"recipient_department_id"`
	SupplierName            string         `db:"supplier_name" json:"supplier_name,omitempty"`
	PurchaseRequestNumber   string         `db:"purchase_request_number" json:"purchase_request_number,omitempty"`
	RecipientDepartmentName string         `db:"recipient_department_name" json:"recipient_department_name,omitempty"`
	Items                   []DeliveryInvoiceItem `json:"items,omitempty"`
}

type DeliveryInvoiceItem struct {
	ID                int     `db:"id" json:"id"`
	DeliveryInvoiceID int     `db:"delivery_invoice_id" json:"delivery_invoice_id"`
	ProductID         int     `db:"product_id" json:"product_id"`
	Quantity          float64 `db:"quantity" json:"quantity"`
	Price             float64 `db:"price" json:"price"`
	ProductName       string  `db:"product_name" json:"product_name,omitempty"`
	UnitName          string  `db:"unit_name" json:"unit_name,omitempty"`
}

type InternalTransferInvoice struct {
	ID                 int        `db:"id" json:"id"`
	TransferDate       string     `db:"transfer_date" json:"transfer_date"`
	FromDepartmentID   int        `db:"from_department_id" json:"from_department_id"`
	ToDepartmentID     int        `db:"to_department_id" json:"to_department_id"`
	Status             string     `db:"status" json:"status"`
	CreatedAt          time.Time  `db:"created_at" json:"created_at"`
	CancelledAt        *time.Time `db:"cancelled_at" json:"cancelled_at,omitempty"`
	FromDepartmentName string     `db:"from_department_name" json:"from_department_name,omitempty"`
	ToDepartmentName   string     `db:"to_department_name" json:"to_department_name,omitempty"`
	Items              []InternalTransferInvoiceItem `json:"items,omitempty"`
}

type InternalTransferInvoiceItem struct {
	ID                        int     `db:"id" json:"id"`
	InternalTransferInvoiceID int     `db:"internal_transfer_invoice_id" json:"internal_transfer_invoice_id"`
	ProductID                 int     `db:"product_id" json:"product_id"`
	Quantity                  float64 `db:"quantity" json:"quantity"`
	ProductName               string  `db:"product_name" json:"product_name,omitempty"`
	UnitName                  string  `db:"unit_name" json:"unit_name,omitempty"`
}

type InventoryBalance struct {
	ID             int       `db:"id" json:"id"`
	DepartmentID   int       `db:"department_id" json:"department_id"`
	ProductID      int       `db:"product_id" json:"product_id"`
	Quantity       float64   `db:"quantity" json:"quantity"`
	UpdatedAt      time.Time `db:"updated_at" json:"updated_at"`
	DepartmentName string    `db:"department_name" json:"department_name,omitempty"`
	ProductName    string    `db:"product_name" json:"product_name,omitempty"`
	SKU            string    `db:"sku" json:"sku,omitempty"`
	UnitName       string    `db:"unit_name" json:"unit_name,omitempty"`
}

type APIResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
	Meta    *Meta       `json:"meta,omitempty"`
}

type Meta struct {
	Total      int `json:"total,omitempty"`
	Page       int `json:"page,omitempty"`
	PerPage    int `json:"per_page,omitempty"`
	TotalPages int `json:"total_pages,omitempty"`
}

type ListFilter struct {
	Page     int    `json:"page"`
	PerPage  int    `json:"per_page"`
	Search   string `json:"search"`
	SortBy   string `json:"sort_by"`
	SortDir  string `json:"sort_dir"`
	Status   string `json:"status"`
	DateFrom string `json:"date_from"`
	DateTo   string `json:"date_to"`
}

func (f *ListFilter) Normalize() {
	if f.Page < 1 {
		f.Page = 1
	}
	if f.PerPage < 1 {
		f.PerPage = 20
	}
	if f.PerPage > 100 {
		f.PerPage = 100
	}
	if f.SortDir != "asc" && f.SortDir != "desc" {
		f.SortDir = "desc"
	}
}

func NullTime(t time.Time) sql.NullTime {
	return sql.NullTime{Time: t, Valid: !t.IsZero()}
}

func NullString(s string) sql.NullString {
	return sql.NullString{String: s, Valid: s != ""}
}

func NullInt64(i int) sql.NullInt64 {
	return sql.NullInt64{Int64: int64(i), Valid: i != 0}
}

func NullFloat64(f float64) sql.NullFloat64 {
	return sql.NullFloat64{Float64: f, Valid: f != 0}
}
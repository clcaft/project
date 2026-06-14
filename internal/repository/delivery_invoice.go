package repository

import (
	"database/sql"
	"fmt"
	"procurement-system/internal/model"
)

type DeliveryInvoiceRepository interface {
	GetAll(filter model.ListFilter) ([]model.DeliveryInvoice, int, error)
	GetByID(id int) (*model.DeliveryInvoice, error)
	Create(di *model.DeliveryInvoice) error
	Update(di *model.DeliveryInvoice) error
	Delete(id int) error
	ChangeStatus(id int, status string) error
}

type deliveryInvoiceRepository struct {
	db *sql.DB
}

func NewDeliveryInvoiceRepository(db *sql.DB) DeliveryInvoiceRepository {
	return &deliveryInvoiceRepository{db: db}
}

func (r *deliveryInvoiceRepository) GetAll(filter model.ListFilter) ([]model.DeliveryInvoice, int, error) {
	query := `
		SELECT di.id, di.invoice_date, di.supplier_id, di.purchase_request_id, di.status,
		       di.created_at, di.cancelled_at, di.recipient_department_id,
		       s.name as supplier_name, d.name as recipient_department_name
		FROM delivery_invoices di
		LEFT JOIN suppliers s ON di.supplier_id = s.id
		LEFT JOIN departments d ON di.recipient_department_id = d.id
		ORDER BY di.id DESC
		LIMIT ? OFFSET ?`

	rows, err := r.db.Query(query, filter.PerPage, (filter.Page-1)*filter.PerPage)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to query delivery invoices: %w", err)
	}
	defer rows.Close()

	var invoices []model.DeliveryInvoice
	for rows.Next() {
		var di model.DeliveryInvoice
		var prID sql.NullInt64
		var cancelledAt sql.NullTime

		if err := rows.Scan(&di.ID, &di.InvoiceDate, &di.SupplierID, &prID, &di.Status,
			&di.CreatedAt, &cancelledAt, &di.RecipientDepartmentID,
			&di.SupplierName, &di.RecipientDepartmentName); err != nil {
			return nil, 0, err
		}
		if prID.Valid {
			pid := int(prID.Int64)
			di.PurchaseRequestID = &pid
		}
		if cancelledAt.Valid {
			di.CancelledAt = &cancelledAt.Time
		}
		invoices = append(invoices, di)
	}

	var total int
	r.db.QueryRow("SELECT COUNT(*) FROM delivery_invoices").Scan(&total)
	return invoices, total, nil
}

func (r *deliveryInvoiceRepository) GetByID(id int) (*model.DeliveryInvoice, error) {
	var di model.DeliveryInvoice
	var prID sql.NullInt64
	var cancelledAt sql.NullTime

	err := r.db.QueryRow(`
		SELECT di.id, di.invoice_date, di.supplier_id, di.purchase_request_id, di.status,
		       di.created_at, di.cancelled_at, di.recipient_department_id,
		       s.name as supplier_name, d.name as recipient_department_name
		FROM delivery_invoices di
		LEFT JOIN suppliers s ON di.supplier_id = s.id
		LEFT JOIN departments d ON di.recipient_department_id = d.id
		WHERE di.id = ?`, id,
	).Scan(&di.ID, &di.InvoiceDate, &di.SupplierID, &prID, &di.Status,
		&di.CreatedAt, &cancelledAt, &di.RecipientDepartmentID,
		&di.SupplierName, &di.RecipientDepartmentName)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get delivery invoice: %w", err)
	}
	if prID.Valid {
		pid := int(prID.Int64)
		di.PurchaseRequestID = &pid
	}
	if cancelledAt.Valid {
		di.CancelledAt = &cancelledAt.Time
	}

	items, err := r.getItems(id)
	if err != nil {
		return nil, err
	}
	di.Items = items
	return &di, nil
}

func (r *deliveryInvoiceRepository) getItems(diID int) ([]model.DeliveryInvoiceItem, error) {
	rows, err := r.db.Query(`
		SELECT dii.id, dii.delivery_invoice_id, dii.product_id, dii.quantity, dii.price,
		       p.name as product_name, pu.short_name as unit_name
		FROM delivery_invoice_items dii
		JOIN products p ON dii.product_id = p.id
		JOIN product_units pu ON p.unit_id = pu.id
		WHERE dii.delivery_invoice_id = ?`, diID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []model.DeliveryInvoiceItem
	for rows.Next() {
		var item model.DeliveryInvoiceItem
		if err := rows.Scan(&item.ID, &item.DeliveryInvoiceID, &item.ProductID, &item.Quantity, &item.Price,
			&item.ProductName, &item.UnitName); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (r *deliveryInvoiceRepository) Create(di *model.DeliveryInvoice) error {
	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	result, err := tx.Exec(
		"INSERT INTO delivery_invoices (invoice_date, supplier_id, purchase_request_id, status, recipient_department_id) VALUES (?, ?, ?, ?, ?)",
		di.InvoiceDate, di.SupplierID, di.PurchaseRequestID, di.Status, di.RecipientDepartmentID)
	if err != nil {
		return fmt.Errorf("failed to create delivery invoice: %w", err)
	}
	id, _ := result.LastInsertId()
	di.ID = int(id)

	for i := range di.Items {
		_, err := tx.Exec(
			"INSERT INTO delivery_invoice_items (delivery_invoice_id, product_id, quantity, price) VALUES (?, ?, ?, ?)",
			di.ID, di.Items[i].ProductID, di.Items[i].Quantity, di.Items[i].Price)
		if err != nil {
			return fmt.Errorf("failed to create delivery invoice item: %w", err)
		}
	}

	return tx.Commit()
}

func (r *deliveryInvoiceRepository) Update(di *model.DeliveryInvoice) error {
	_, err := r.db.Exec(
		"UPDATE delivery_invoices SET invoice_date = ?, supplier_id = ?, purchase_request_id = ?, recipient_department_id = ? WHERE id = ?",
		di.InvoiceDate, di.SupplierID, di.PurchaseRequestID, di.RecipientDepartmentID, di.ID)
	if err != nil {
		return fmt.Errorf("failed to update delivery invoice: %w", err)
	}
	return nil
}

func (r *deliveryInvoiceRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM delivery_invoices WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("failed to delete delivery invoice: %w", err)
	}
	return nil
}

func (r *deliveryInvoiceRepository) ChangeStatus(id int, status string) error {
	_, err := r.db.Exec("UPDATE delivery_invoices SET status = ? WHERE id = ?", status, id)
	if err != nil {
		return fmt.Errorf("failed to change status: %w", err)
	}
	return nil
}

package repository

import (
	"database/sql"
	"fmt"
	"procurement-system/internal/model"
)

type PurchaseRequestRepository interface {
	GetAll(filter model.ListFilter) ([]model.PurchaseRequest, int, error)
	GetByID(id int) (*model.PurchaseRequest, error)
	Create(pr *model.PurchaseRequest) error
	Update(pr *model.PurchaseRequest) error
	Delete(id int) error
	ChangeStatus(id int, status string) error
}

type purchaseRequestRepository struct {
	db *sql.DB
}

func NewPurchaseRequestRepository(db *sql.DB) PurchaseRequestRepository {
	return &purchaseRequestRepository{db: db}
}

func (r *purchaseRequestRepository) GetAll(filter model.ListFilter) ([]model.PurchaseRequest, int, error) {
	query := `
		SELECT pr.id, pr.request_date, pr.planned_delivery_date, pr.supplier_id, pr.status,
		       pr.created_at, pr.cancelled_at, pr.recipient_department_id,
		       s.name as supplier_name, d.name as recipient_department_name
		FROM purchase_requests pr
		LEFT JOIN suppliers s ON pr.supplier_id = s.id
		LEFT JOIN departments d ON pr.recipient_department_id = d.id
		ORDER BY pr.id DESC
		LIMIT ? OFFSET ?`

	rows, err := r.db.Query(query, filter.PerPage, (filter.Page-1)*filter.PerPage)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to query purchase requests: %w", err)
	}
	defer rows.Close()

	var requests []model.PurchaseRequest
	for rows.Next() {
		var pr model.PurchaseRequest
		var supplierID sql.NullInt64
		var plannedDate sql.NullString
		var cancelledAt sql.NullTime

		if err := rows.Scan(&pr.ID, &pr.RequestDate, &plannedDate, &supplierID, &pr.Status,
			&pr.CreatedAt, &cancelledAt, &pr.RecipientDepartmentID,
			&pr.SupplierName, &pr.RecipientDepartmentName); err != nil {
			return nil, 0, err
		}
		if supplierID.Valid {
			sid := int(supplierID.Int64)
			pr.SupplierID = &sid
		}
		if plannedDate.Valid {
			pr.PlannedDeliveryDate = &plannedDate.String
		}
		if cancelledAt.Valid {
			pr.CancelledAt = &cancelledAt.Time
		}
		requests = append(requests, pr)
	}

	var total int
	r.db.QueryRow("SELECT COUNT(*) FROM purchase_requests").Scan(&total)
	return requests, total, nil
}

func (r *purchaseRequestRepository) GetByID(id int) (*model.PurchaseRequest, error) {
	var pr model.PurchaseRequest
	var supplierID sql.NullInt64
	var plannedDate sql.NullString
	var cancelledAt sql.NullTime

	err := r.db.QueryRow(`
		SELECT pr.id, pr.request_date, pr.planned_delivery_date, pr.supplier_id, pr.status,
		       pr.created_at, pr.cancelled_at, pr.recipient_department_id,
		       s.name as supplier_name, d.name as recipient_department_name
		FROM purchase_requests pr
		LEFT JOIN suppliers s ON pr.supplier_id = s.id
		LEFT JOIN departments d ON pr.recipient_department_id = d.id
		WHERE pr.id = ?`, id,
	).Scan(&pr.ID, &pr.RequestDate, &plannedDate, &supplierID, &pr.Status,
		&pr.CreatedAt, &cancelledAt, &pr.RecipientDepartmentID,
		&pr.SupplierName, &pr.RecipientDepartmentName)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get purchase request: %w", err)
	}
	if supplierID.Valid {
		sid := int(supplierID.Int64)
		pr.SupplierID = &sid
	}
	if plannedDate.Valid {
		pr.PlannedDeliveryDate = &plannedDate.String
	}
	if cancelledAt.Valid {
		pr.CancelledAt = &cancelledAt.Time
	}

	items, err := r.getItems(id)
	if err != nil {
		return nil, err
	}
	pr.Items = items
	return &pr, nil
}

func (r *purchaseRequestRepository) getItems(prID int) ([]model.PurchaseRequestItem, error) {
	rows, err := r.db.Query(`
		SELECT pri.id, pri.purchase_request_id, pri.product_id, pri.quantity,
		       p.name as product_name, pu.short_name as unit_name
		FROM purchase_request_items pri
		JOIN products p ON pri.product_id = p.id
		JOIN product_units pu ON p.unit_id = pu.id
		WHERE pri.purchase_request_id = ?`, prID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []model.PurchaseRequestItem
	for rows.Next() {
		var item model.PurchaseRequestItem
		if err := rows.Scan(&item.ID, &item.PurchaseRequestID, &item.ProductID, &item.Quantity,
			&item.ProductName, &item.UnitName); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (r *purchaseRequestRepository) Create(pr *model.PurchaseRequest) error {
	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	result, err := tx.Exec(
		"INSERT INTO purchase_requests (request_date, planned_delivery_date, supplier_id, status, recipient_department_id) VALUES (?, ?, ?, ?, ?)",
		pr.RequestDate, pr.PlannedDeliveryDate, pr.SupplierID, pr.Status, pr.RecipientDepartmentID)
	if err != nil {
		return fmt.Errorf("failed to create purchase request: %w", err)
	}
	id, _ := result.LastInsertId()
	pr.ID = int(id)

	for i := range pr.Items {
		_, err := tx.Exec(
			"INSERT INTO purchase_request_items (purchase_request_id, product_id, quantity) VALUES (?, ?, ?)",
			pr.ID, pr.Items[i].ProductID, pr.Items[i].Quantity)
		if err != nil {
			return fmt.Errorf("failed to create purchase request item: %w", err)
		}
	}
	return tx.Commit()
}

func (r *purchaseRequestRepository) Update(pr *model.PurchaseRequest) error {
	_, err := r.db.Exec(
		"UPDATE purchase_requests SET request_date = ?, planned_delivery_date = ?, supplier_id = ?, recipient_department_id = ? WHERE id = ?",
		pr.RequestDate, pr.PlannedDeliveryDate, pr.SupplierID, pr.RecipientDepartmentID, pr.ID)
	if err != nil {
		return fmt.Errorf("failed to update purchase request: %w", err)
	}
	return nil
}

func (r *purchaseRequestRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM purchase_requests WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("failed to delete purchase request: %w", err)
	}
	return nil
}

func (r *purchaseRequestRepository) ChangeStatus(id int, status string) error {
	_, err := r.db.Exec("UPDATE purchase_requests SET status = ? WHERE id = ?", status, id)
	if err != nil {
		return fmt.Errorf("failed to change status: %w", err)
	}
	return nil
}

package repository

import (
	"database/sql"
	"fmt"
	"procurement-system/internal/model"
)

type InternalTransferInvoiceRepository interface {
	GetAll(filter model.ListFilter) ([]model.InternalTransferInvoice, int, error)
	GetByID(id int) (*model.InternalTransferInvoice, error)
	Create(it *model.InternalTransferInvoice) error
	Update(it *model.InternalTransferInvoice) error
	Delete(id int) error
	ChangeStatus(id int, status string) error
}

type internalTransferInvoiceRepository struct {
	db *sql.DB
}

func NewInternalTransferInvoiceRepository(db *sql.DB) InternalTransferInvoiceRepository {
	return &internalTransferInvoiceRepository{db: db}
}

func (r *internalTransferInvoiceRepository) GetAll(filter model.ListFilter) ([]model.InternalTransferInvoice, int, error) {
	baseQuery := `
		SELECT iti.id, iti.transfer_date, iti.from_department_id, iti.to_department_id, iti.status,
<<<<<<< HEAD
		       iti.created_at, iti.cancelled_at,
=======
		       iti.created_at, iti.cancelled_at, iti.created_by, iti.notes,
>>>>>>> fc07f468f8ab1a3e8bbde8aad30dcf077a584766
		       fd.name as from_department_name, td.name as to_department_name
		FROM internal_transfer_invoices iti
		LEFT JOIN departments fd ON iti.from_department_id = fd.id
		LEFT JOIN departments td ON iti.to_department_id = td.id`

	var conditions []string
	var args []interface{}

	if filter.Status != "" {
		conditions = append(conditions, "iti.status = ?")
		args = append(args, filter.Status)
	}
	if filter.DateFrom != "" {
		conditions = append(conditions, "iti.transfer_date >= ?")
		args = append(args, filter.DateFrom)
	}
	if filter.DateTo != "" {
		conditions = append(conditions, "iti.transfer_date <= ?")
		args = append(args, filter.DateTo)
	}

	query := baseQuery
	if len(conditions) > 0 {
		query += " WHERE " + conditions[0]
		for _, c := range conditions[1:] {
			query += " AND " + c
		}
	}
	query += " ORDER BY iti.id DESC"

	offset := (filter.Page - 1) * filter.PerPage
	query += fmt.Sprintf(" LIMIT %d OFFSET %d", filter.PerPage, offset)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to query transfer invoices: %w", err)
	}
	defer rows.Close()

	var invoices []model.InternalTransferInvoice
	for rows.Next() {
		var it model.InternalTransferInvoice
<<<<<<< HEAD
		var cancelledAt sql.NullTime

		if err := rows.Scan(&it.ID, &it.TransferDate, &it.FromDepartmentID, &it.ToDepartmentID, &it.Status,
			&it.CreatedAt, &cancelledAt,
			&it.FromDepartmentName, &it.ToDepartmentName); err != nil {
			return nil, 0, err
		}
=======
		var createdBy sql.NullInt64
		var cancelledAt sql.NullTime

		if err := rows.Scan(&it.ID, &it.TransferDate, &it.FromDepartmentID, &it.ToDepartmentID, &it.Status,
			&it.CreatedAt, &cancelledAt, &createdBy, &it.Notes,
			&it.FromDepartmentName, &it.ToDepartmentName); err != nil {
			return nil, 0, err
		}
		if createdBy.Valid {
			cb := int(createdBy.Int64)
			it.CreatedBy = &cb
		}
>>>>>>> fc07f468f8ab1a3e8bbde8aad30dcf077a584766
		if cancelledAt.Valid {
			it.CancelledAt = &cancelledAt.Time
		}
		invoices = append(invoices, it)
	}

	var total int
	err = r.db.QueryRow(`SELECT COUNT(*) FROM internal_transfer_invoices`).Scan(&total)
	return invoices, total, err
}

func (r *internalTransferInvoiceRepository) GetByID(id int) (*model.InternalTransferInvoice, error) {
	var it model.InternalTransferInvoice
<<<<<<< HEAD
=======
	var createdBy sql.NullInt64
>>>>>>> fc07f468f8ab1a3e8bbde8aad30dcf077a584766
	var cancelledAt sql.NullTime

	err := r.db.QueryRow(`
		SELECT iti.id, iti.transfer_date, iti.from_department_id, iti.to_department_id, iti.status,
<<<<<<< HEAD
		       iti.created_at, iti.cancelled_at,
=======
		       iti.created_at, iti.cancelled_at, iti.created_by, iti.notes,
>>>>>>> fc07f468f8ab1a3e8bbde8aad30dcf077a584766
		       fd.name as from_department_name, td.name as to_department_name
		FROM internal_transfer_invoices iti
		LEFT JOIN departments fd ON iti.from_department_id = fd.id
		LEFT JOIN departments td ON iti.to_department_id = td.id
		WHERE iti.id = ?`, id,
	).Scan(&it.ID, &it.TransferDate, &it.FromDepartmentID, &it.ToDepartmentID, &it.Status,
<<<<<<< HEAD
		&it.CreatedAt, &cancelledAt,
=======
		&it.CreatedAt, &cancelledAt, &createdBy, &it.Notes,
>>>>>>> fc07f468f8ab1a3e8bbde8aad30dcf077a584766
		&it.FromDepartmentName, &it.ToDepartmentName)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get transfer invoice: %w", err)
	}
<<<<<<< HEAD
=======
	if createdBy.Valid {
		cb := int(createdBy.Int64)
		it.CreatedBy = &cb
	}
>>>>>>> fc07f468f8ab1a3e8bbde8aad30dcf077a584766
	if cancelledAt.Valid {
		it.CancelledAt = &cancelledAt.Time
	}

	items, err := r.getItems(id)
	if err != nil {
		return nil, err
	}
	it.Items = items

	return &it, nil
}

func (r *internalTransferInvoiceRepository) getItems(itID int) ([]model.InternalTransferInvoiceItem, error) {
	rows, err := r.db.Query(`
<<<<<<< HEAD
		SELECT itii.id, itii.internal_transfer_invoice_id, itii.product_id, itii.quantity,
=======
		SELECT itii.id, itii.internal_transfer_invoice_id, itii.product_id, itii.quantity, itii.notes, itii.created_at,
>>>>>>> fc07f468f8ab1a3e8bbde8aad30dcf077a584766
		       p.name as product_name, pu.short_name as unit_name
		FROM internal_transfer_invoice_items itii
		JOIN products p ON itii.product_id = p.id
		JOIN product_units pu ON p.unit_id = pu.id
		WHERE itii.internal_transfer_invoice_id = ?`, itID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []model.InternalTransferInvoiceItem
	for rows.Next() {
		var item model.InternalTransferInvoiceItem
		if err := rows.Scan(&item.ID, &item.InternalTransferInvoiceID, &item.ProductID, &item.Quantity,
<<<<<<< HEAD
			&item.ProductName, &item.UnitName); err != nil {
=======
			&item.Notes, &item.CreatedAt, &item.ProductName, &item.UnitName); err != nil {
>>>>>>> fc07f468f8ab1a3e8bbde8aad30dcf077a584766
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (r *internalTransferInvoiceRepository) Create(it *model.InternalTransferInvoice) error {
	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	result, err := tx.Exec(
<<<<<<< HEAD
		`INSERT INTO internal_transfer_invoices (transfer_date, from_department_id, to_department_id, status)
		 VALUES (?, ?, ?, ?)`,
		it.TransferDate, it.FromDepartmentID, it.ToDepartmentID, it.Status,
=======
		`INSERT INTO internal_transfer_invoices (transfer_date, from_department_id, to_department_id, status, notes)
		 VALUES (?, ?, ?, ?, ?)`,
		it.TransferDate, it.FromDepartmentID, it.ToDepartmentID, it.Status, it.Notes,
>>>>>>> fc07f468f8ab1a3e8bbde8aad30dcf077a584766
	)
	if err != nil {
		return fmt.Errorf("failed to create transfer invoice: %w", err)
	}
	id, _ := result.LastInsertId()
	it.ID = int(id)

	for i := range it.Items {
		_, err := tx.Exec(
<<<<<<< HEAD
			`INSERT INTO internal_transfer_invoice_items (internal_transfer_invoice_id, product_id, quantity)
			 VALUES (?, ?, ?)`,
			it.ID, it.Items[i].ProductID, it.Items[i].Quantity,
=======
			`INSERT INTO internal_transfer_invoice_items (internal_transfer_invoice_id, product_id, quantity, notes)
			 VALUES (?, ?, ?, ?)`,
			it.ID, it.Items[i].ProductID, it.Items[i].Quantity, it.Items[i].Notes,
>>>>>>> fc07f468f8ab1a3e8bbde8aad30dcf077a584766
		)
		if err != nil {
			return fmt.Errorf("failed to create transfer invoice item: %w", err)
		}
	}

	return tx.Commit()
}

func (r *internalTransferInvoiceRepository) Update(it *model.InternalTransferInvoice) error {
	_, err := r.db.Exec(
<<<<<<< HEAD
		`UPDATE internal_transfer_invoices SET transfer_date = ?, from_department_id = ?, to_department_id = ? WHERE id = ?`,
		it.TransferDate, it.FromDepartmentID, it.ToDepartmentID, it.ID,
=======
		`UPDATE internal_transfer_invoices SET transfer_date = ?, from_department_id = ?, to_department_id = ?, notes = ? WHERE id = ?`,
		it.TransferDate, it.FromDepartmentID, it.ToDepartmentID, it.Notes, it.ID,
>>>>>>> fc07f468f8ab1a3e8bbde8aad30dcf077a584766
	)
	if err != nil {
		return fmt.Errorf("failed to update transfer invoice: %w", err)
	}
	return nil
}

func (r *internalTransferInvoiceRepository) Delete(id int) error {
	_, err := r.db.Exec(`DELETE FROM internal_transfer_invoices WHERE id = ?`, id)
	if err != nil {
		return fmt.Errorf("failed to delete transfer invoice: %w", err)
	}
	return nil
}

func (r *internalTransferInvoiceRepository) ChangeStatus(id int, status string) error {
	_, err := r.db.Exec(`UPDATE internal_transfer_invoices SET status = ? WHERE id = ?`, status, id)
	if err != nil {
		return fmt.Errorf("failed to change status: %w", err)
	}
	return nil
}
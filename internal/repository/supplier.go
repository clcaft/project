package repository

import (
	"database/sql"
	"fmt"
	"procurement-system/internal/model"
)

type SupplierRepository interface {
	GetAll(filter model.ListFilter) ([]model.Supplier, int, error)
	GetByID(id int) (*model.Supplier, error)
	Create(s *model.Supplier) error
	Update(s *model.Supplier) error
	Delete(id int) error
}

type supplierRepository struct {
	db *sql.DB
}

func NewSupplierRepository(db *sql.DB) SupplierRepository {
	return &supplierRepository{db: db}
}

func (r *supplierRepository) GetAll(filter model.ListFilter) ([]model.Supplier, int, error) {
	allowedSort := map[string]string{"id": "id", "name": "name"}
	query, args := buildListQuery("SELECT id, name, phone, email, is_active, address, created_at, updated_at FROM suppliers", filter, allowedSort)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to query suppliers: %w", err)
	}
	defer rows.Close()

	var suppliers []model.Supplier
	for rows.Next() {
		var s model.Supplier
		if err := rows.Scan(&s.ID, &s.Name, &s.Phone, &s.Email, &s.IsActive, &s.Address, &s.CreatedAt, &s.UpdatedAt); err != nil {
			return nil, 0, err
		}
		suppliers = append(suppliers, s)
	}

	total, err := countQuery(r.db, "SELECT COUNT(*) FROM suppliers", filter)
	return suppliers, total, err
}

func (r *supplierRepository) GetByID(id int) (*model.Supplier, error) {
	var s model.Supplier
	err := r.db.QueryRow("SELECT id, name, phone, email, is_active, address, created_at, updated_at FROM suppliers WHERE id = ?", id,
	).Scan(&s.ID, &s.Name, &s.Phone, &s.Email, &s.IsActive, &s.Address, &s.CreatedAt, &s.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get supplier: %w", err)
	}
	return &s, nil
}

func (r *supplierRepository) Create(s *model.Supplier) error {
	result, err := r.db.Exec("INSERT INTO suppliers (name, phone, email, is_active, address) VALUES (?, ?, ?, ?, ?)", s.Name, s.Phone, s.Email, s.IsActive, s.Address)
	if err != nil {
		return fmt.Errorf("failed to create supplier: %w", err)
	}
	id, _ := result.LastInsertId()
	s.ID = int(id)
	return nil
}

func (r *supplierRepository) Update(s *model.Supplier) error {
	_, err := r.db.Exec("UPDATE suppliers SET name = ?, phone = ?, email = ?, is_active = ?, address = ? WHERE id = ?", s.Name, s.Phone, s.Email, s.IsActive, s.Address, s.ID)
	if err != nil {
		return fmt.Errorf("failed to update supplier: %w", err)
	}
	return nil
}

func (r *supplierRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM suppliers WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("failed to delete supplier: %w", err)
	}
	return nil
}

package repository

import (
	"database/sql"
	"fmt"
	"procurement-system/internal/model"
)

type WarehouseRepository interface {
	GetAll(filter model.ListFilter) ([]model.Warehouse, int, error)
	GetByID(id int) (*model.Warehouse, error)
	Create(w *model.Warehouse) error
	Update(w *model.Warehouse) error
	Delete(id int) error
}

type warehouseRepository struct{ db *sql.DB }

func NewWarehouseRepository(db *sql.DB) WarehouseRepository { return &warehouseRepository{db: db} }

func (r *warehouseRepository) GetAll(filter model.ListFilter) ([]model.Warehouse, int, error) {
	rows, err := r.db.Query(`
<<<<<<< HEAD
		SELECT w.id, w.address, w.warehouse_type, w.department_id,
=======
		SELECT w.id, w.address, w.warehouse_type, w.department_id, w.is_active, w.created_at, w.updated_at,
>>>>>>> fc07f468f8ab1a3e8bbde8aad30dcf077a584766
		       d.name as department_name
		FROM warehouses w
		JOIN departments d ON w.department_id = d.id
		ORDER BY w.id DESC`)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var items []model.Warehouse
	for rows.Next() {
		var w model.Warehouse
<<<<<<< HEAD
		if err := rows.Scan(&w.ID, &w.Address, &w.WarehouseType, &w.DepartmentID); err != nil {
=======
		if err := rows.Scan(&w.ID, &w.Address, &w.WarehouseType, &w.DepartmentID, &w.IsActive, &w.CreatedAt, &w.UpdatedAt); err != nil {
>>>>>>> fc07f468f8ab1a3e8bbde8aad30dcf077a584766
			return nil, 0, err
		}
		items = append(items, w)
	}
	var total int
	r.db.QueryRow(`SELECT COUNT(*) FROM warehouses`).Scan(&total)
	return items, total, nil
}

func (r *warehouseRepository) GetByID(id int) (*model.Warehouse, error) {
	var w model.Warehouse
	err := r.db.QueryRow(`
<<<<<<< HEAD
		SELECT id, address, warehouse_type, department_id
		FROM warehouses WHERE id = ?`, id,
	).Scan(&w.ID, &w.Address, &w.WarehouseType, &w.DepartmentID)
=======
		SELECT id, address, warehouse_type, department_id, is_active, created_at, updated_at
		FROM warehouses WHERE id = ?`, id,
	).Scan(&w.ID, &w.Address, &w.WarehouseType, &w.DepartmentID, &w.IsActive, &w.CreatedAt, &w.UpdatedAt)
>>>>>>> fc07f468f8ab1a3e8bbde8aad30dcf077a584766
	if err == sql.ErrNoRows { return nil, nil }
	if err != nil { return nil, fmt.Errorf("failed to get warehouse: %w", err) }
	return &w, nil
}

func (r *warehouseRepository) Create(w *model.Warehouse) error {
	result, err := r.db.Exec(
<<<<<<< HEAD
		`INSERT INTO warehouses (address, warehouse_type, department_id) VALUES (?, ?, ?)`,
		w.Address, w.WarehouseType, w.DepartmentID,
=======
		`INSERT INTO warehouses (address, warehouse_type, department_id, is_active) VALUES (?, ?, ?, ?)`,
		w.Address, w.WarehouseType, w.DepartmentID, w.IsActive,
>>>>>>> fc07f468f8ab1a3e8bbde8aad30dcf077a584766
	)
	if err != nil { return fmt.Errorf("failed to create warehouse: %w", err) }
	id, _ := result.LastInsertId()
	w.ID = int(id)
	return nil
}

func (r *warehouseRepository) Update(w *model.Warehouse) error {
	_, err := r.db.Exec(
<<<<<<< HEAD
		`UPDATE warehouses SET address = ?, warehouse_type = ?, department_id = ? WHERE id = ?`,
		w.Address, w.WarehouseType, w.DepartmentID, w.ID,
=======
		`UPDATE warehouses SET address = ?, warehouse_type = ?, department_id = ?, is_active = ? WHERE id = ?`,
		w.Address, w.WarehouseType, w.DepartmentID, w.IsActive, w.ID,
>>>>>>> fc07f468f8ab1a3e8bbde8aad30dcf077a584766
	)
	if err != nil { return fmt.Errorf("failed to update warehouse: %w", err) }
	return nil
}

func (r *warehouseRepository) Delete(id int) error {
	_, err := r.db.Exec(`DELETE FROM warehouses WHERE id = ?`, id)
	if err != nil { return fmt.Errorf("failed to delete warehouse: %w", err) }
	return nil
}
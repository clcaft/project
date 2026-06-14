package repository

import (
	"database/sql"
	"fmt"
	"procurement-system/internal/model"
)

type StoreRepository interface {
	GetAll(filter model.ListFilter) ([]model.Store, int, error)
	GetByID(id int) (*model.Store, error)
	Create(s *model.Store) error
	Update(s *model.Store) error
	Delete(id int) error
}

type storeRepository struct{ db *sql.DB }

func NewStoreRepository(db *sql.DB) StoreRepository { return &storeRepository{db: db} }

func (r *storeRepository) GetAll(filter model.ListFilter) ([]model.Store, int, error) {
	rows, err := r.db.Query(`
<<<<<<< HEAD
		SELECT id, address, store_type, department_id
=======
		SELECT id, address, store_type, department_id, is_active, created_at, updated_at
>>>>>>> fc07f468f8ab1a3e8bbde8aad30dcf077a584766
		FROM stores ORDER BY id DESC`)
	if err != nil { return nil, 0, err }
	defer rows.Close()

	var items []model.Store
	for rows.Next() {
		var s model.Store
<<<<<<< HEAD
		if err := rows.Scan(&s.ID, &s.Address, &s.StoreType, &s.DepartmentID); err != nil {
=======
		if err := rows.Scan(&s.ID, &s.Address, &s.StoreType, &s.DepartmentID, &s.IsActive, &s.CreatedAt, &s.UpdatedAt); err != nil {
>>>>>>> fc07f468f8ab1a3e8bbde8aad30dcf077a584766
			return nil, 0, err
		}
		items = append(items, s)
	}
	var total int
	r.db.QueryRow(`SELECT COUNT(*) FROM stores`).Scan(&total)
	return items, total, nil
}

func (r *storeRepository) GetByID(id int) (*model.Store, error) {
	var s model.Store
<<<<<<< HEAD
	err := r.db.QueryRow(`SELECT id, address, store_type, department_id FROM stores WHERE id = ?`, id,
	).Scan(&s.ID, &s.Address, &s.StoreType, &s.DepartmentID)
=======
	err := r.db.QueryRow(`SELECT id, address, store_type, department_id, is_active, created_at, updated_at FROM stores WHERE id = ?`, id,
	).Scan(&s.ID, &s.Address, &s.StoreType, &s.DepartmentID, &s.IsActive, &s.CreatedAt, &s.UpdatedAt)
>>>>>>> fc07f468f8ab1a3e8bbde8aad30dcf077a584766
	if err == sql.ErrNoRows { return nil, nil }
	if err != nil { return nil, fmt.Errorf("failed to get store: %w", err) }
	return &s, nil
}

func (r *storeRepository) Create(s *model.Store) error {
<<<<<<< HEAD
	result, err := r.db.Exec(`INSERT INTO stores (address, store_type, department_id) VALUES (?, ?, ?)`,
		s.Address, s.StoreType, s.DepartmentID)
=======
	result, err := r.db.Exec(`INSERT INTO stores (address, store_type, department_id, is_active) VALUES (?, ?, ?, ?)`,
		s.Address, s.StoreType, s.DepartmentID, s.IsActive)
>>>>>>> fc07f468f8ab1a3e8bbde8aad30dcf077a584766
	if err != nil { return fmt.Errorf("failed to create store: %w", err) }
	id, _ := result.LastInsertId()
	s.ID = int(id)
	return nil
}

func (r *storeRepository) Update(s *model.Store) error {
<<<<<<< HEAD
	_, err := r.db.Exec(`UPDATE stores SET address = ?, store_type = ?, department_id = ? WHERE id = ?`,
		s.Address, s.StoreType, s.DepartmentID, s.ID)
=======
	_, err := r.db.Exec(`UPDATE stores SET address = ?, store_type = ?, department_id = ?, is_active = ? WHERE id = ?`,
		s.Address, s.StoreType, s.DepartmentID, s.IsActive, s.ID)
>>>>>>> fc07f468f8ab1a3e8bbde8aad30dcf077a584766
	if err != nil { return fmt.Errorf("failed to update store: %w", err) }
	return nil
}

func (r *storeRepository) Delete(id int) error {
	_, err := r.db.Exec(`DELETE FROM stores WHERE id = ?`, id)
	if err != nil { return fmt.Errorf("failed to delete store: %w", err) }
	return nil
}
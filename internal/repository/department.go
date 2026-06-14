package repository

import (
	"database/sql"
	"fmt"
	"procurement-system/internal/model"
)

type DepartmentRepository interface {
	GetAll(filter model.ListFilter) ([]model.Department, int, error)
	GetByID(id int) (*model.Department, error)
	Create(d *model.Department) error
	Update(d *model.Department) error
	Delete(id int) error
}

type departmentRepository struct {
	db *sql.DB
}

func NewDepartmentRepository(db *sql.DB) DepartmentRepository {
	return &departmentRepository{db: db}
}

func (r *departmentRepository) GetAll(filter model.ListFilter) ([]model.Department, int, error) {
	allowedSort := map[string]string{"id": "id", "name": "name", "type": "department_type"}
	query, args := buildListQuery("SELECT id, name, department_type, is_active FROM departments", filter, allowedSort)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to query departments: %w", err)
	}
	defer rows.Close()

	var departments []model.Department
	for rows.Next() {
		var d model.Department
		if err := rows.Scan(&d.ID, &d.Name, &d.DepartmentType, &d.IsActive); err != nil {
			return nil, 0, err
		}
		departments = append(departments, d)
	}

	total, err := countQuery(r.db, "SELECT COUNT(*) FROM departments", filter)
	return departments, total, err
}

func (r *departmentRepository) GetByID(id int) (*model.Department, error) {
	var d model.Department
	err := r.db.QueryRow("SELECT id, name, department_type, is_active FROM departments WHERE id = ?", id,
	).Scan(&d.ID, &d.Name, &d.DepartmentType, &d.IsActive)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get department: %w", err)
	}
	return &d, nil
}

func (r *departmentRepository) Create(d *model.Department) error {
	result, err := r.db.Exec("INSERT INTO departments (name, department_type, is_active) VALUES (?, ?, ?)", d.Name, d.DepartmentType, d.IsActive)
	if err != nil {
		return fmt.Errorf("failed to create department: %w", err)
	}
	id, _ := result.LastInsertId()
	d.ID = int(id)
	return nil
}

func (r *departmentRepository) Update(d *model.Department) error {
	_, err := r.db.Exec("UPDATE departments SET name = ?, department_type = ?, is_active = ? WHERE id = ?", d.Name, d.DepartmentType, d.IsActive, d.ID)
	if err != nil {
		return fmt.Errorf("failed to update department: %w", err)
	}
	return nil
}

func (r *departmentRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM departments WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("failed to delete department: %w", err)
	}
	return nil
}
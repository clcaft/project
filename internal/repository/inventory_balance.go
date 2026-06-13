package repository

import (
	"database/sql"
	"fmt"
	"procurement-system/internal/model"
)

type InventoryBalanceRepository interface {
	GetAll(filter model.ListFilter) ([]model.InventoryBalance, int, error)
	GetByDepartmentAndProduct(departmentID, productID int) (*model.InventoryBalance, error)
	GetByDepartment(departmentID int) ([]model.InventoryBalance, error)
}

type inventoryBalanceRepository struct {
	db *sql.DB
}

func NewInventoryBalanceRepository(db *sql.DB) InventoryBalanceRepository {
	return &inventoryBalanceRepository{db: db}
}

func (r *inventoryBalanceRepository) GetAll(filter model.ListFilter) ([]model.InventoryBalance, int, error) {
	baseQuery := `
		SELECT ib.id, ib.department_id, ib.product_id, ib.quantity, ib.avg_price, ib.total_amount, ib.updated_at,
		       d.name as department_name, p.name as product_name, p.sku, pu.short_name as unit_name
		FROM inventory_balances ib
		JOIN departments d ON ib.department_id = d.id
		JOIN products p ON ib.product_id = p.id
		JOIN product_units pu ON p.unit_id = pu.id`

	var conditions []string
	var args []interface{}

	if filter.Search != "" {
		conditions = append(conditions, "(p.name LIKE ? OR p.sku LIKE ?)")
		args = append(args, "%"+filter.Search+"%", "%"+filter.Search+"%")
	}

	query := baseQuery
	if len(conditions) > 0 {
		query += " WHERE " + conditions[0]
		for _, c := range conditions[1:] {
			query += " AND " + c
		}
	}
	query += " ORDER BY ib.updated_at DESC"

	offset := (filter.Page - 1) * filter.PerPage
	query += fmt.Sprintf(" LIMIT %d OFFSET %d", filter.PerPage, offset)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to query inventory balances: %w", err)
	}
	defer rows.Close()

	var balances []model.InventoryBalance
	for rows.Next() {
		var b model.InventoryBalance
		if err := rows.Scan(&b.ID, &b.DepartmentID, &b.ProductID, &b.Quantity, &b.AvgPrice, &b.TotalAmount, &b.UpdatedAt,
			&b.DepartmentName, &b.ProductName, &b.SKU, &b.UnitName); err != nil {
			return nil, 0, err
		}
		balances = append(balances, b)
	}

	var total int
	err = r.db.QueryRow(`SELECT COUNT(*) FROM inventory_balances`).Scan(&total)
	return balances, total, err
}

func (r *inventoryBalanceRepository) GetByDepartmentAndProduct(departmentID, productID int) (*model.InventoryBalance, error) {
	var b model.InventoryBalance
	err := r.db.QueryRow(`
		SELECT ib.id, ib.department_id, ib.product_id, ib.quantity, ib.avg_price, ib.total_amount, ib.updated_at,
		       d.name as department_name, p.name as product_name, p.sku, pu.short_name as unit_name
		FROM inventory_balances ib
		JOIN departments d ON ib.department_id = d.id
		JOIN products p ON ib.product_id = p.id
		JOIN product_units pu ON p.unit_id = pu.id
		WHERE ib.department_id = ? AND ib.product_id = ?`, departmentID, productID,
	).Scan(&b.ID, &b.DepartmentID, &b.ProductID, &b.Quantity, &b.AvgPrice, &b.TotalAmount, &b.UpdatedAt,
		&b.DepartmentName, &b.ProductName, &b.SKU, &b.UnitName)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get inventory balance: %w", err)
	}
	return &b, nil
}

func (r *inventoryBalanceRepository) GetByDepartment(departmentID int) ([]model.InventoryBalance, error) {
	rows, err := r.db.Query(`
		SELECT ib.id, ib.department_id, ib.product_id, ib.quantity, ib.avg_price, ib.total_amount, ib.updated_at,
		       d.name as department_name, p.name as product_name, p.sku, pu.short_name as unit_name
		FROM inventory_balances ib
		JOIN departments d ON ib.department_id = d.id
		JOIN products p ON ib.product_id = p.id
		JOIN product_units pu ON p.unit_id = pu.id
		WHERE ib.department_id = ?
		ORDER BY p.name`, departmentID)
	if err != nil {
		return nil, fmt.Errorf("failed to query department balances: %w", err)
	}
	defer rows.Close()

	var balances []model.InventoryBalance
	for rows.Next() {
		var b model.InventoryBalance
		if err := rows.Scan(&b.ID, &b.DepartmentID, &b.ProductID, &b.Quantity, &b.AvgPrice, &b.TotalAmount, &b.UpdatedAt,
			&b.DepartmentName, &b.ProductName, &b.SKU, &b.UnitName); err != nil {
			return nil, err
		}
		balances = append(balances, b)
	}
	return balances, nil
}
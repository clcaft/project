package repository

import (
	"database/sql"
	"fmt"
	"procurement-system/internal/model"
)

type ProductUnitRepository interface {
	GetAll(filter model.ListFilter) ([]model.ProductUnit, int, error)
	GetByID(id int) (*model.ProductUnit, error)
	Create(pu *model.ProductUnit) error
	Update(pu *model.ProductUnit) error
	Delete(id int) error
}

type productUnitRepository struct{ db *sql.DB }

func NewProductUnitRepository(db *sql.DB) ProductUnitRepository { return &productUnitRepository{db: db} }

func (r *productUnitRepository) GetAll(filter model.ListFilter) ([]model.ProductUnit, int, error) {
	rows, err := r.db.Query(`SELECT id, name, short_name FROM product_units ORDER BY id DESC`)
	if err != nil { return nil, 0, err }
	defer rows.Close()

	var items []model.ProductUnit
	for rows.Next() {
		var pu model.ProductUnit
		if err := rows.Scan(&pu.ID, &pu.Name, &pu.ShortName); err != nil {
			return nil, 0, err
		}
		items = append(items, pu)
	}
	var total int
	r.db.QueryRow(`SELECT COUNT(*) FROM product_units`).Scan(&total)
	return items, total, nil
}

func (r *productUnitRepository) GetByID(id int) (*model.ProductUnit, error) {
	var pu model.ProductUnit
	err := r.db.QueryRow(`SELECT id, name, short_name FROM product_units WHERE id = ?`, id,
	).Scan(&pu.ID, &pu.Name, &pu.ShortName)
	if err == sql.ErrNoRows { return nil, nil }
	if err != nil { return nil, fmt.Errorf("failed to get unit: %w", err) }
	return &pu, nil
}

func (r *productUnitRepository) Create(pu *model.ProductUnit) error {
	result, err := r.db.Exec(`INSERT INTO product_units (name, short_name) VALUES (?, ?)`, pu.Name, pu.ShortName)
	if err != nil { return fmt.Errorf("failed to create unit: %w", err) }
	id, _ := result.LastInsertId()
	pu.ID = int(id)
	return nil
}

func (r *productUnitRepository) Update(pu *model.ProductUnit) error {
	_, err := r.db.Exec(`UPDATE product_units SET name = ?, short_name = ? WHERE id = ?`, pu.Name, pu.ShortName, pu.ID)
	if err != nil { return fmt.Errorf("failed to update unit: %w", err) }
	return nil
}

func (r *productUnitRepository) Delete(id int) error {
	_, err := r.db.Exec(`DELETE FROM product_units WHERE id = ?`, id)
	if err != nil { return fmt.Errorf("failed to delete unit: %w", err) }
	return nil
}
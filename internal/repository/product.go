package repository

import (
	"database/sql"
	"fmt"
	"procurement-system/internal/model"
)

type ProductRepository interface {
	GetAll(filter model.ListFilter) ([]model.Product, int, error)
	GetByID(id int) (*model.Product, error)
	Create(p *model.Product) error
	Update(p *model.Product) error
	Delete(id int) error
}

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) GetAll(filter model.ListFilter) ([]model.Product, int, error) {
	query := `
		SELECT p.id, p.name, p.sku, p.unit_id, p.category_id, p.is_active,
		       pu.short_name as unit_name, pc.name as category_name
		FROM products p
		LEFT JOIN product_units pu ON p.unit_id = pu.id
		LEFT JOIN product_categories pc ON p.category_id = pc.id
		ORDER BY p.id DESC
		LIMIT ? OFFSET ?`

	rows, err := r.db.Query(query, filter.PerPage, (filter.Page-1)*filter.PerPage)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to query products: %w", err)
	}
	defer rows.Close()

	var products []model.Product
	for rows.Next() {
		var p model.Product
		var catID sql.NullInt64
		if err := rows.Scan(&p.ID, &p.Name, &p.SKU, &p.UnitID, &catID, &p.IsActive, &p.UnitName, &p.CategoryName); err != nil {
			return nil, 0, err
		}
		if catID.Valid {
			cid := int(catID.Int64)
			p.CategoryID = &cid
		}
		products = append(products, p)
	}

	var total int
	r.db.QueryRow("SELECT COUNT(*) FROM products").Scan(&total)
	return products, total, nil
}

func (r *productRepository) GetByID(id int) (*model.Product, error) {
	var p model.Product
	var catID sql.NullInt64
	err := r.db.QueryRow(`
		SELECT p.id, p.name, p.sku, p.unit_id, p.category_id, p.is_active,
		       pu.short_name as unit_name, pc.name as category_name
		FROM products p
		LEFT JOIN product_units pu ON p.unit_id = pu.id
		LEFT JOIN product_categories pc ON p.category_id = pc.id
		WHERE p.id = ?`, id,
	).Scan(&p.ID, &p.Name, &p.SKU, &p.UnitID, &catID, &p.IsActive, &p.UnitName, &p.CategoryName)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get product: %w", err)
	}
	if catID.Valid {
		cid := int(catID.Int64)
		p.CategoryID = &cid
	}
	return &p, nil
}

func (r *productRepository) Create(p *model.Product) error {
	result, err := r.db.Exec("INSERT INTO products (name, sku, unit_id, category_id, is_active) VALUES (?, ?, ?, ?, ?)", p.Name, p.SKU, p.UnitID, p.CategoryID, p.IsActive)
	if err != nil {
		return fmt.Errorf("failed to create product: %w", err)
	}
	id, _ := result.LastInsertId()
	p.ID = int(id)
	return nil
}

func (r *productRepository) Update(p *model.Product) error {
	_, err := r.db.Exec("UPDATE products SET name = ?, sku = ?, unit_id = ?, category_id = ?, is_active = ? WHERE id = ?", p.Name, p.SKU, p.UnitID, p.CategoryID, p.IsActive, p.ID)
	if err != nil {
		return fmt.Errorf("failed to update product: %w", err)
	}
	return nil
}

func (r *productRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM products WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("failed to delete product: %w", err)
	}
	return nil
}
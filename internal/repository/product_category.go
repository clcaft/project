package repository

import (
	"database/sql"
	"fmt"
	"procurement-system/internal/model"
)

type ProductCategoryRepository interface {
	GetAll(filter model.ListFilter) ([]model.ProductCategory, int, error)
	GetByID(id int) (*model.ProductCategory, error)
	Create(pc *model.ProductCategory) error
	Update(pc *model.ProductCategory) error
	Delete(id int) error
}

type productCategoryRepository struct{ db *sql.DB }

func NewProductCategoryRepository(db *sql.DB) ProductCategoryRepository { return &productCategoryRepository{db: db} }

func (r *productCategoryRepository) GetAll(filter model.ListFilter) ([]model.ProductCategory, int, error) {
<<<<<<< HEAD
	rows, err := r.db.Query(`SELECT id, name FROM product_categories ORDER BY id DESC`)
=======
	rows, err := r.db.Query(`SELECT id, name, description, created_at FROM product_categories ORDER BY id DESC`)
>>>>>>> fc07f468f8ab1a3e8bbde8aad30dcf077a584766
	if err != nil { return nil, 0, err }
	defer rows.Close()

	var items []model.ProductCategory
	for rows.Next() {
		var pc model.ProductCategory
<<<<<<< HEAD
		if err := rows.Scan(&pc.ID, &pc.Name); err != nil {
=======
		if err := rows.Scan(&pc.ID, &pc.Name, &pc.Description, &pc.CreatedAt); err != nil {
>>>>>>> fc07f468f8ab1a3e8bbde8aad30dcf077a584766
			return nil, 0, err
		}
		items = append(items, pc)
	}
	var total int
	r.db.QueryRow(`SELECT COUNT(*) FROM product_categories`).Scan(&total)
	return items, total, nil
}

func (r *productCategoryRepository) GetByID(id int) (*model.ProductCategory, error) {
	var pc model.ProductCategory
<<<<<<< HEAD
	err := r.db.QueryRow(`SELECT id, name FROM product_categories WHERE id = ?`, id,
	).Scan(&pc.ID, &pc.Name)
=======
	err := r.db.QueryRow(`SELECT id, name, description, created_at FROM product_categories WHERE id = ?`, id,
	).Scan(&pc.ID, &pc.Name, &pc.Description, &pc.CreatedAt)
>>>>>>> fc07f468f8ab1a3e8bbde8aad30dcf077a584766
	if err == sql.ErrNoRows { return nil, nil }
	if err != nil { return nil, fmt.Errorf("failed to get category: %w", err) }
	return &pc, nil
}

func (r *productCategoryRepository) Create(pc *model.ProductCategory) error {
<<<<<<< HEAD
	result, err := r.db.Exec(`INSERT INTO product_categories (name) VALUES (?)`, pc.Name)
=======
	result, err := r.db.Exec(`INSERT INTO product_categories (name, description) VALUES (?, ?)`, pc.Name, pc.Description)
>>>>>>> fc07f468f8ab1a3e8bbde8aad30dcf077a584766
	if err != nil { return fmt.Errorf("failed to create category: %w", err) }
	id, _ := result.LastInsertId()
	pc.ID = int(id)
	return nil
}

func (r *productCategoryRepository) Update(pc *model.ProductCategory) error {
<<<<<<< HEAD
	_, err := r.db.Exec(`UPDATE product_categories SET name = ? WHERE id = ?`, pc.Name, pc.ID)
=======
	_, err := r.db.Exec(`UPDATE product_categories SET name = ?, description = ? WHERE id = ?`, pc.Name, pc.Description, pc.ID)
>>>>>>> fc07f468f8ab1a3e8bbde8aad30dcf077a584766
	if err != nil { return fmt.Errorf("failed to update category: %w", err) }
	return nil
}

func (r *productCategoryRepository) Delete(id int) error {
	_, err := r.db.Exec(`DELETE FROM product_categories WHERE id = ?`, id)
	if err != nil { return fmt.Errorf("failed to delete category: %w", err) }
	return nil
}
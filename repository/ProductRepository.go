package repository

import (
	"CrudApp/delivery/Products/dtos"
	"CrudApp/domain"
	"database/sql"
	"github.com/jmoiron/sqlx"
)

type ProductRepo struct {
	db *sqlx.DB
}

func NewProductRepo(db *sqlx.DB) domain.ProductRepository {
	return &ProductRepo{db: db}
}

func (r *ProductRepo) Create(dto dtos.CreateProductDto) (domain.Product, error) {
	var product domain.Product
	query := `INSERT INTO products (name, price) VALUES ($1, $2) RETURNING id, name, price, created_at`
	err := r.db.QueryRowx(query, dto.Name, dto.Price).StructScan(&product)
	return product, err
}

func (r *ProductRepo) GetByID(id int) (domain.Product, error) {
	var product domain.Product
	query := `SELECT id, name, price, created_at FROM products WHERE id = $1`
	err := r.db.Get(&product, query, id)
	return product, err
}

func (r *ProductRepo) Update(id int, dto dtos.UpdateProductDto) (domain.Product, error) {
	var product domain.Product
	query := `UPDATE products SET name = $1, price = $2 WHERE id = $3 RETURNING id, name, price, created_at`
	err := r.db.QueryRowx(query, dto.Name, dto.Price, id).StructScan(&product)
	if err == sql.ErrNoRows {
		return domain.Product{}, sql.ErrNoRows
	}
	return product, err
}

func (r *ProductRepo) GetAll() ([]domain.Product, error) {
	var products []domain.Product
	query := `SELECT id, name, price, created_at FROM products`
	err := r.db.Select(&products, query)
	return products, err
}

func (r *ProductRepo) Delete(id int) error {
	query := `DELETE FROM products WHERE id = $1`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return sql.ErrNoRows
	}
	return nil
}
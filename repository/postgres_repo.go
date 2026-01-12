package repository

import (
	"database/sql"
	"CrudApp/domain"
	_ "github.com/lib/pq"
)

type postgresRepo struct {
	db *sql.DB
}

func NewPostgresRepo(db *sql.DB) domain.ProductRepository {
	return &postgresRepo{db: db}
}

func (r *postgresRepo) Create(p domain.Product) (domain.Product, error) {
	err := r.db.QueryRow("INSERT INTO products (name, price) VALUES ($1, $2) RETURNING id", p.Name, p.Price).Scan(&p.ID)
	return p, err
}

func (r *postgresRepo) GetAll() ([]domain.Product, error) {
	rows, err := r.db.Query("SELECT id, name, price FROM products")
	if err != nil { return nil, err }
	defer rows.Close()

	var products []domain.Product
	for rows.Next() {
		var p domain.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price); err != nil { return nil, err }
		products = append(products, p)
	}
	return products, nil
}

func (r *postgresRepo) GetByID(id int) (domain.Product, error) {
	var p domain.Product
	err := r.db.QueryRow("SELECT id, name, price FROM products WHERE id = $1", id).Scan(&p.ID, &p.Name, &p.Price)
	return p, err
}

func (r *postgresRepo) Update(id int, p domain.Product) (domain.Product, error) {
	_, err := r.db.Exec("UPDATE products SET name = $1, price = $2 WHERE id = $3", p.Name, p.Price, id)
	p.ID = id
	return p, err
}

func (r *postgresRepo) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM products WHERE id = $1", id)
	return err
}
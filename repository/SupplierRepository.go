package repository

import (
	"CrudApp/delivery/Suppliers/dtos"
	"CrudApp/domain"
	"github.com/jmoiron/sqlx"
)

type SupplierRepo struct {
	db *sqlx.DB
}

func NewSupplierRepo(db *sqlx.DB) domain.SupplierRepository {
	return &SupplierRepo{db: db}
}

func (r *SupplierRepo) Create(dto dtos.CreateSupplierDto) (dtos.CreateSupplierDto, error) {
	query := `INSERT INTO suppliers (id, name, contact) VALUES ($1, $2, $3)`
	_, err := r.db.Exec(query, dto.Id, dto.Name, dto.Contact)
	return dto, err
}

func (r *SupplierRepo) GetAll() ([]domain.Supplier, error) {
	var suppliers []domain.Supplier
	query := `SELECT id, name, contact FROM suppliers`
	err := r.db.Select(&suppliers, query)
	return suppliers, err
}
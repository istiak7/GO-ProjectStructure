package repository

import (
	"CrudApp/domain"
	"gorm.io/gorm"
)

type ProductRepo struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) domain.ProductRepository {
	return &ProductRepo{db: db}
}

func (r *ProductRepo) Create(p domain.Product) (domain.Product, error) {
	result := r.db.Create(&p)
	return p, result.Error
}

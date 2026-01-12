package repository

import (
	"gorm.io/gorm"
	"CrudApp/domain"
	_ "github.com/lib/pq"
)

type postgresRepo struct {
	db *gorm.DB
}

func NewPostgresRepo(db *gorm.DB) domain.ProductRepository {
	return &postgresRepo{db: db}
}

func (r *postgresRepo) Create(p domain.Product) (domain.Product, error) {
	result := r.db.Create(&p)
	return p, result.Error
}

func (r *postgresRepo) GetAll() ([]domain.Product, error) {
	var products []domain.Product
	result := r.db.Find(&products)
	return products, result.Error
}

func (r *postgresRepo) GetByID(id int) (domain.Product, error) {
	var p domain.Product
	result := r.db.First(&p, id)
	return p, result.Error
}

func (r *postgresRepo) Update(id int, p domain.Product) (domain.Product, error) {
	p.ID = id
	result := r.db.Save(&p)
	return p, result.Error
}

func (r *postgresRepo) Delete(id int) error {
	result := r.db.Delete(&domain.Product{}, id)
	return result.Error
}
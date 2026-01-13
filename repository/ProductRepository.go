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

func (r *ProductRepo) GetByID(id int) (domain.Product, error) {
	var product domain.Product
	result := r.db.First(&product, id)
	return product, result.Error
}

func (r *ProductRepo) Update(id int, p domain.Product) (domain.Product, error) {
	var product domain.Product
	result := r.db.First(&product, id)
	if result.Error != nil {
		return domain.Product{}, result.Error
	}
	result = r.db.Model(&product).Updates(p)
	return product, result.Error
}

func (r *ProductRepo) GetAll() ([]domain.Product, error) {
	var products []domain.Product
	result := r.db.Find(&products)
	return products, result.Error
}

func (r *ProductRepo) Delete(id int) error {
	result := r.db.Delete(&domain.Product{}, id)
	return result.Error
}
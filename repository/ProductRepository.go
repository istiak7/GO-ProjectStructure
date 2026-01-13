package repository

import (
	"CrudApp/delivery/Products/dtos"
	"CrudApp/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProductRepo struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) domain.ProductRepository {
	return &ProductRepo{db: db}
}

func (r *ProductRepo) Create(dto dtos.CreateProductDto) (domain.Product, error) {
	p := domain.Product{
		Name:    dto.Name,
		Price:   dto.Price,

	}
	result := r.db.Create(&p)
	return p, result.Error
}

func (r *ProductRepo) GetByID(id int) (domain.Product, error) {
	var product domain.Product
	result := r.db.First(&product, id)
	return product, result.Error
}

func (r *ProductRepo) Update(id int, dto dtos.UpdateProductDto) (domain.Product, error) {
	
	product := domain.Product{
		Name:  dto.Name,
		Price: dto.Price,
	}

	result := r.db.Model(&product).
		Clauses(clause.Returning{}). 
		Where("id = ?", id).
		Select("*").       
		Updates(product)

	if result.Error != nil {
		return domain.Product{}, result.Error
	}

	if result.RowsAffected == 0 {
		return domain.Product{}, gorm.ErrRecordNotFound
	}

	return product, nil
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
package repository

import (
	"CrudApp/delivery/Suppliers/dtos"
	"CrudApp/domain"
	"gorm.io/gorm"
)

type SupplierRepo struct {
	db *gorm.DB
}

func NewSupplierRepo(db *gorm.DB) domain.SupplierRepository {
	return &SupplierRepo{db: db}
}

func (r *SupplierRepo) Create(dto dtos.CreateSupplierDto) (dtos.CreateSupplierDto, error) {
	s := domain.Supplier{
		ID:      dto.Id,
		Name:    dto.Name,
		Contact: dto.Contact,
	}

	result := r.db.Create(&s)
	return dto, result.Error
}

func (r *SupplierRepo) GetAll() ([]domain.Supplier, error) {
	var suppliers []domain.Supplier
	result := r.db.Find(&suppliers)
	return suppliers, result.Error
}
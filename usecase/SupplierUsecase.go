package usecase

import (
	"CrudApp/delivery/Suppliers/dtos"
	"CrudApp/domain"
)

type SupplierUsecase struct {
	Repo domain.SupplierRepository
}

func (u *SupplierUsecase) CreateSupplier(dto dtos.CreateSupplierDto) (dtos.CreateSupplierDto, error) {
	return u.Repo.Create(dto)
}
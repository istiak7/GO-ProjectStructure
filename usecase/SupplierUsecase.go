package usecase
import (
	"CrudApp/domain"
	 Dtos "CrudApp/delivery/Suppliers/Dtos"
)
type SupplierUsecase struct {
	Repo domain.SupplierRepository
}

func (u * SupplierUsecase) CreateSupplier(dto Dtos.CreateSupplierDto) (Dtos.CreateSupplierDto, error) {
	return u.Repo.Create(dto)
}
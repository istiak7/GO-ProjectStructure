package usecase

import (
	"CrudApp/domain"
)

type ProductUseCase struct {
	Repo domain.ProductRepository
}

func (u *ProductUseCase) CreateProduct(p domain.Product) (domain.Product, error) {
	return u.Repo.Create(p)
}

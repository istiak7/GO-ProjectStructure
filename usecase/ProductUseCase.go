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

func (u *ProductUseCase) GetProductByID(id int) (domain.Product, error) {
	return u.Repo.GetByID(id)
}

func (u *ProductUseCase) UpdateProduct(id int, p domain.Product) (domain.Product, error) {
	return u.Repo.Update(id, p)
}

func (u *ProductUseCase) DeleteProduct(id int) error {
	return u.Repo.Delete(id)
}
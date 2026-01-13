package usecase

import (
	"CrudApp/domain"
	"CrudApp/delivery/Products/dtos"
)

type ProductUseCase struct {
	Repo domain.ProductRepository
}

func (u *ProductUseCase) CreateProduct(dto dtos.CreateProductDto) (domain.Product, error) {
	return u.Repo.Create(dto)
}

func (u *ProductUseCase) GetProductByID(id int) (domain.Product, error) {
	return u.Repo.GetByID(id)
}

func (u *ProductUseCase) UpdateProduct(id int, dto dtos.UpdateProductDto) (domain.Product, error) {
	return u.Repo.Update(id, dto)
}

func (u *ProductUseCase) DeleteProduct(id int) error {
	return u.Repo.Delete(id)
}
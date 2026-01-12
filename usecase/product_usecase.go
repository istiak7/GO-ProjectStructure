package usecase

import "CrudApp/domain"

type ProductUsecase struct {
	Repo domain.ProductRepository
}

func (u *ProductUsecase) CreateProduct(p domain.Product) (domain.Product, error) { return u.Repo.Create(p) }
func (u *ProductUsecase) ListProducts() ([]domain.Product, error)           { return u.Repo.GetAll() }
func (u *ProductUsecase) GetProduct(id int) (domain.Product, error)         { return u.Repo.GetByID(id) }
func (u *ProductUsecase) UpdateProduct(id int, p domain.Product) (domain.Product, error) { return u.Repo.Update(id, p) }
func (u *ProductUsecase) RemoveProduct(id int) error                        { return u.Repo.Delete(id) }
package config

import (
	"CrudApp/delivery/Products"
	"CrudApp/delivery/Suppliers"
	"CrudApp/repository"
	"CrudApp/route"
	"CrudApp/usecase"
	"github.com/jmoiron/sqlx"
)

func InitDependencies(db *sqlx.DB) *route.Handlers {
	supplierRepo := repository.NewSupplierRepo(db)
	supplierUC := &usecase.SupplierUsecase{Repo: supplierRepo}
	supplierHandler := &Suppliers.SupplierHandler{Usecase: supplierUC}

	productRepo := repository.NewProductRepo(db)
	productUC := &usecase.ProductUseCase{Repo: productRepo}
	productHandler := &products.ProductHandler{Usecase: productUC}
    
	return &route.Handlers{
		Supplier: supplierHandler,
		Product:  productHandler,
	}
}

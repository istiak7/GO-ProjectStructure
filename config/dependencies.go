package config

import (
	"CrudApp/delivery/Suppliers"
	"CrudApp/repository"
	"CrudApp/route"
	"CrudApp/usecase"
	"gorm.io/gorm"
)

func InitDependencies(db *gorm.DB) *route.Handlers {
	supplierRepo := repository.NewSupplierRepo(db)
	supplierUC := &usecase.SupplierUsecase{Repo: supplierRepo}
	supplierHandler := &Suppliers.SupplierHandler{Usecase: supplierUC}
    
	return &route.Handlers{
		Supplier: supplierHandler,
		// Product: productHandler,
	}
}

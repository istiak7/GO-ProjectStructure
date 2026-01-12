package config

import (
	"CrudApp/delivery/Suppliers"
	"CrudApp/repository"
	"CrudApp/usecase"
	"gorm.io/gorm"
)

func InitDependencies(db *gorm.DB) *Suppliers.SupplierHandler {
	supplierRepo := repository.NewSupplierRepo(db)
	supplierUC := &usecase.SupplierUsecase{Repo: supplierRepo}
	return &Suppliers.SupplierHandler{Usecase: supplierUC}
}

package domain

import (
	"CrudApp/delivery/Suppliers/dtos"
	"time"
)

type Supplier struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Contact   string    `json:"contact"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type SupplierRepository interface {
	Create(dtos.CreateSupplierDto) (dtos.CreateSupplierDto, error)
}
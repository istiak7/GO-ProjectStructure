package domain

import (
	Dtos "CrudApp/delivery/Suppliers/Dtos"
	"time"
)

type Supplier struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Contact   string    `json:"contact"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type SupplierRepository interface {
	Create(Dtos.CreateSupplierDto) (Dtos.CreateSupplierDto, error)
}

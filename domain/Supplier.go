package domain

import (
	"CrudApp/delivery/Suppliers/dtos"
	"time"
)

type Supplier struct {
	ID        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Contact   string    `json:"contact" db:"contact"`
	Address   string    `json:"address" db:"address"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type SupplierRepository interface {
	Create(dtos.CreateSupplierDto) (dtos.CreateSupplierDto, error)
}
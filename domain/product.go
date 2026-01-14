package domain

import (
	"CrudApp/delivery/Products/dtos"
	"time"
)

type Product struct {
	ID        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Price     float64   `json:"price" db:"price"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type ProductRepository interface {
	Create(dto dtos.CreateProductDto) (Product, error)
	GetByID(id int) (Product, error)	
	Update(id int, dto dtos.UpdateProductDto) (Product, error)
	Delete(id int) error
}

package domain

import (
	"CrudApp/delivery/Products/dtos"
	"time"
)

type Product struct {
	ID        int       `json:"id" example:"1" gorm:"primaryKey"`
	Name      string    `json:"name" example:"Gaming Keyboard"`
	Price     float64   `json:"price" example:"75.50"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}

type ProductRepository interface {
	Create(dto dtos.CreateProductDto) (Product, error)
	GetByID(id int) (Product, error)	
	Update(id int, dto dtos.UpdateProductDto) (Product, error)
	Delete(id int) error
}

package dtos
type CreateProductDto struct {

	Name string `json:"name" validate:"required,min=2,max=100"`
	Price float64 `json:"price" validate:"required,gt=0"`
}
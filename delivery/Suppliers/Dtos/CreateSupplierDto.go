package dtos

type CreateSupplierDto struct {
	Id  int    `json:"id"`
	Name    string `json:"name"`
	Contact string `json:"contact"`  
}
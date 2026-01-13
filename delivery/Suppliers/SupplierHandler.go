package Suppliers

import (
	Dtos "CrudApp/delivery/Suppliers/Dtos"
	"CrudApp/usecase"
	"encoding/json"
	"net/http"
)

type SupplierHandler struct {
	Usecase *usecase.SupplierUsecase
}


func (h *SupplierHandler) CreateSupplier(w http.ResponseWriter, r *http.Request) {

	var dto Dtos.CreateSupplierDto
	json.NewDecoder(r.Body).Decode(&dto)
	res, _ := h.Usecase.CreateSupplier(dto)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)

}

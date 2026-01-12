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

// CreateSupplier godoc
// @Summary Create a new supplier
// @Description Create a new supplier
// @Tags suppliers
// @Accept json
// @Produce json
// @Param supplier body dtos.CreateSupplierDto true "Supplier data"
// @Success 201 {object} map[string]interface{}
// @Router /suppliers [post]
func (h *SupplierHandler) CreateSupplier(w http.ResponseWriter, r *http.Request) {

	var dto Dtos.CreateSupplierDto
	json.NewDecoder(r.Body).Decode(&dto)
	res, _ := h.Usecase.CreateSupplier(dto)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)

}

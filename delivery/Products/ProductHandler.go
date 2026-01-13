package products

import (
	"CrudApp/domain"
	"CrudApp/usecase"
	"encoding/json"
	"net/http"
)

type ProductHandler struct {
	Usecase *usecase.ProductUseCase
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product domain.Product
	json.NewDecoder(r.Body).Decode(&product)
	res, err := h.Usecase.CreateProduct(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

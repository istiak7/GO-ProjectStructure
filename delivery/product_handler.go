package delivery

import (
	"encoding/json"
	"CrudApp/domain"
	"CrudApp/usecase"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	Usecase *usecase.ProductUsecase
}

// @Summary List all products
// @Tags products
// @Success 200 {array} domain.Product
// @Router /products [get]
func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	res, _ := h.Usecase.ListProducts()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

// @Summary Create a product
// @Tags products
// @Param product body domain.Product true "Product Body"
// @Success 201 {object} domain.Product
// @Router /products [post]
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var p domain.Product
	json.NewDecoder(r.Body).Decode(&p)
	res, _ := h.Usecase.CreateProduct(p)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

// @Summary Update a product
// @Tags products
// @Param id path int true "Product ID"
// @Param product body domain.Product true "New Data"
// @Router /products/{id} [put]
func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.PathValue("id")) // Go 1.22+
	var p domain.Product
	json.NewDecoder(r.Body).Decode(&p)
	res, _ := h.Usecase.UpdateProduct(id, p)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

// @Summary Delete a product
// @Tags products
// @Param id path int true "Product ID"
// @Router /products/{id} [delete]
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.PathValue("id"))
	h.Usecase.RemoveProduct(id)
	w.WriteHeader(http.StatusNoContent)
}
package products

import (
	"CrudApp/common/utilities"
	"CrudApp/delivery/Products/dtos"
	"CrudApp/usecase"
	"encoding/json"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	Usecase *usecase.ProductUseCase
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var dto dtos.CreateProductDto

	// Json decode & format check 
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		utilities.SendResponse(w, http.StatusBadRequest, false, "Invalid JSON format", nil, err.Error())
		return
	}

	// Validation check
	if validationErrors := utilities.ValidateStruct(dto); len(validationErrors) > 0 {
		utilities.SendResponse(w, http.StatusBadRequest, false, "Validation failed", validationErrors, nil)
		return
	}

	// Usecase call
	res, err := h.Usecase.CreateProduct(dto)
	if err != nil {
		utilities.SendResponse(w, http.StatusInternalServerError, false, "Failed to create product", nil, err.Error())
		return
	}

	//Success
	utilities.SendResponse(w, http.StatusCreated, true, "Product created successfully", res, nil)
}

func (h *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utilities.SendResponse(w, http.StatusBadRequest, false, "Invalid ID format", nil, "ID must be a number")
		return
	}

	product, err := h.Usecase.GetProductByID(id)
	if err != nil {
	
		utilities.SendResponse(w, http.StatusNotFound, false, "Product not found", nil, err.Error())
		return
	}


	utilities.SendResponse(w, http.StatusOK, true, "Product fetched successfully", product, nil)
}

func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utilities.SendResponse(w, http.StatusBadRequest, false, "Invalid ID format", nil, "ID must be a number")
		return
	}

	
	var dto dtos.UpdateProductDto
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		utilities.SendResponse(w, http.StatusBadRequest, false, "Invalid JSON format", nil, err.Error())
		return
	}

	if validationErrors := utilities.ValidateStruct(dto); len(validationErrors) > 0 {
		utilities.SendResponse(w, http.StatusBadRequest, false, "Validation failed", validationErrors, nil)
		return
	}

	res, err := h.Usecase.UpdateProduct(id, dto)
	if err != nil {
		utilities.SendResponse(w, http.StatusInternalServerError, false, "Failed to update product", nil, err.Error())
		return
	}

	utilities.SendResponse(w, http.StatusOK, true, "Product updated successfully", res, nil)
}

func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utilities.SendResponse(w, http.StatusBadRequest, false, "Invalid ID format", nil, "ID must be a number")
		return
	}

	err = h.Usecase.DeleteProduct(id)
	if err != nil {
		utilities.SendResponse(w, http.StatusInternalServerError, false, "Failed to delete product", nil, err.Error())
		return
	}

	utilities.SendResponse(w, http.StatusOK, true, "Product deleted successfully", nil, nil)
}
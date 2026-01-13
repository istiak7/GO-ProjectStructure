package route

import (
	"net/http"

	"CrudApp/delivery/Products"
)

func RegisterProductRoutes(mux *http.ServeMux, productHandler *products.ProductHandler) {
	mux.HandleFunc("POST /CreateProducts", productHandler.CreateProduct)
	mux.HandleFunc("GET /GetProducts/{id}", productHandler.GetProductByID)
	mux.HandleFunc("PUT /UpdateProducts/{id}", productHandler.UpdateProduct)
	mux.HandleFunc("DELETE /DeleteProducts/{id}", productHandler.DeleteProduct)
}

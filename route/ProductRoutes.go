package route

import (
	"net/http"

	"CrudApp/delivery/Products"
)

func RegisterProductRoutes(mux *http.ServeMux, productHandler *products.ProductHandler) {
	mux.HandleFunc("POST /products", productHandler.CreateProduct)
	mux.HandleFunc("GET /products/{id}", productHandler.GetProductByID)
	mux.HandleFunc("PUT /products/{id}", productHandler.UpdateProduct)
	mux.HandleFunc("DELETE /products/{id}", productHandler.DeleteProduct)
}

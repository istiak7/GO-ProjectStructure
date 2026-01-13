package route

import (
	"net/http"

	"CrudApp/delivery/Suppliers"
)

type Handlers struct {
	Supplier *Suppliers.SupplierHandler
	// Product  *Products.ProductHandler
	// Auth     *Auth.AuthHandler
}

func SetupRoutes(mux *http.ServeMux, handlers *Handlers) {
	RegisterSwaggerRoutes(mux)
	RegisterSupplierRoutes(mux, handlers.Supplier)
	// RegisterProductRoutes(mux, handlers.Product)
	// RegisterLoginRoutes(mux, handlers.Auth)
}

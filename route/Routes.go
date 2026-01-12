package route

import (
	"net/http"

	"CrudApp/delivery/Suppliers"
)

// SetupRoutes initializes all application routes
func SetupRoutes(mux *http.ServeMux, supplierHandler *Suppliers.SupplierHandler) {
	// Register all route groups

	RegisterSwaggerRoutes(mux)
	RegisterSupplierRoutes(mux, supplierHandler)
	// RegisterProductRoutes(mux, productHandler)
	// RegisterLoginRoutes(mux, authHandler)
	
}

package route

import (
	"net/http"

	"CrudApp/delivery/Suppliers"
)

func RegisterSupplierRoutes(mux *http.ServeMux, supplierHandler *Suppliers.SupplierHandler) {
	mux.HandleFunc("POST /suppliers", supplierHandler.CreateSupplier)
	// Add more supplier routes here as needed
	// mux.HandleFunc("GET /suppliers", supplierHandler.GetSuppliers)
	// mux.HandleFunc("GET /suppliers/{id}", supplierHandler.GetSupplierByID)
	// mux.HandleFunc("PUT /suppliers/{id}", supplierHandler.UpdateSupplier)
	// mux.HandleFunc("DELETE /suppliers/{id}", supplierHandler.DeleteSupplier)
}

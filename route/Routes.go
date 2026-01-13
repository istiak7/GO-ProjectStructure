package route

import (
	"net/http"

	"CrudApp/delivery/Products"
	"CrudApp/delivery/Suppliers"
)

type Handlers struct {
	Supplier *Suppliers.SupplierHandler
	Product  *products.ProductHandler
}

func SetupRoutes(mux *http.ServeMux, handlers *Handlers) {
	RegisterSupplierRoutes(mux, handlers.Supplier)
	RegisterProductRoutes(mux, handlers.Product)
}

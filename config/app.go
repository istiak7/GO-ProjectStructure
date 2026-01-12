package config

import (
	"net/http"

	"CrudApp/route"
)

func InitApp() *http.ServeMux {
	db := InitDB()
	supplierHandler := InitDependencies(db)

	mux := http.NewServeMux()
	route.SetupRoutes(mux, supplierHandler)

	return mux
}

package config

import (
	"net/http"

	"CrudApp/route"
)

func InitApp() *http.ServeMux {
	db := InitDB()
	handlers := InitDependencies(db)

	mux := http.NewServeMux()
	route.SetupRoutes(mux, handlers)

	return mux
}

package config

import (
	"net/http"

	"CrudApp/common/middleware"
	"CrudApp/route"
)

func InitApp() http.Handler {
	db := InitDB()
	handlers := InitDependencies(db)

	mux := http.NewServeMux()
	route.SetupRoutes(mux, handlers)

	return middleware.CORS(mux)
}

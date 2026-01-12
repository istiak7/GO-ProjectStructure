package route

import (
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
)

func RegisterSwaggerRoutes(mux *http.ServeMux) {
	mux.Handle("/swagger/", httpSwagger.WrapHandler)
}

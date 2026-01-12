package main

import (
	"fmt"
	"log"
	"net/http"

	"CrudApp/config"
	_ "CrudApp/docs"
)

// @title Product & Supplier API
// @version 1.0
// @description API for managing products and suppliers
// @host localhost:8080
// @BasePath /

func main() {
	mux := config.InitApp()

	fmt.Println("Server running at :8080")
	fmt.Println("Swagger UI: http://localhost:8080/swagger/index.html")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
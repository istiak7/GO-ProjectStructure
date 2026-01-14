package main

import (
	"fmt"
	"log"
	"net/http"
	"CrudApp/config"
)

func main() {
	handler := config.InitApp()

	fmt.Println("Server running at :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
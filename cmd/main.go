package main

import (
	"fmt"
	"log"
	"net/http"

	"CrudApp/config"
)

func main() {
	mux := config.InitApp()

	fmt.Println("Server running at :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}